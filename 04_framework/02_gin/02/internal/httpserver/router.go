package httpserver

import (
	"auth/internal/config"
	"auth/internal/user"

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

	r.POST("/create",userHanlder.CreateUser)
	r.POST("/login",userHanlder.Login)


	return r
}