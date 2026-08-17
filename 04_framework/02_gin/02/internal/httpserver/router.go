package httpserver

import (
	"auth/internal/config"
	"auth/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(ptrDbPool *pgxpool.Pool,cfg config.Config) *gin.Engine{
	r:=gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health",health)

	userRepo:=user.NewRepository(ptrDbPool)
	userService:=user.NewService(userRepo,cfg.JWT_SECRET)
	userHanlder:=user.NewHandle(userService)

	// public API
	r.POST("/create",userHanlder.CreateUser)
	r.POST("/login",userHanlder.Login)

	// Protected API
	protectedApi:=r.Group("/api")

	protectedApi.GET("/marketwatch",func (c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"ok":true,
			"stock1":"tcs",
			"stock2":"wipro",
			"stock3":"bcl",
		})
	})

	protectedApi.GET("/wallet",func (c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"ok":true,
			"money":"5000",
		})
	})


	return r
}