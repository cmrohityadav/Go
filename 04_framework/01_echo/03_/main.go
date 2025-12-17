package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

var jwtSecret = []byte("mysecret")

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password,omitempty"` //hide in response

}

func main() {
	dns := "host=localhost user=postgres password=123456 dbname=testdb port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db = database

	db.AutoMigrate(&User{})

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.POST("/register", register)
	e.POST("/login",login)

	e.Logger.Fatal(e.Start(":8000"))
}

func register(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(500, echo.Map{"error": "Failed to hash the password"})
	}

	u.Password = string(hash)

	if err := db.Create(&u).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Email already exists"})
	}
	return c.JSON(http.StatusOK, echo.Map{"error": "user registered successfully", "data": u})
}

func login(c echo.Context) error{
	req:=new(User)

	if err:= c.Bind(req); err!=nil{
		return err;
	}
	var user User;
	if err:=db.Where("email=?",req.Email).First(&user).Error;err!=nil{
		return c.JSON(http.StatusUnauthorized,echo.Map{"error":"Invalid email or password"});
	}
	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(req.Password));err!=nil{
		return c.JSON(http.StatusUnauthorized,echo.Map{"error":"Invalid email or password"});
	}
	

	fmt.Println("REQ PASSWORD:", req.Password)
fmt.Println("DB PASSWORD:", user.Password)
	jwtHeaderPayload:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id":user.ID,
		"email":user.Email,
		"exp":time.Now().Add(time.Hour).Unix(),
	})

	jwtToken,err:=jwtHeaderPayload.SignedString([]byte("mySecret"))
	if err!=nil{
		return err;
	}
	return c.JSON(http.StatusOK,echo.Map{
		"message":"Login successfully",
		"token":jwtToken,
	})


}
