package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey=[]byte("heyrohit")


func main(){
	claims:=jwt.MapClaims{
		"sub":"12345",
		"name":"rohit",
		"email":"rohit@gmail.com",
		"exp":time.Now().Add(time.Hour).Unix(),
	}

	//Header + Payload
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	signedToken,err:=token.SignedString(secretKey);
	if err!=nil{
		panic(err)
	}
	fmt.Println("Signed token is given below")
	fmt.Println(signedToken)
	
}