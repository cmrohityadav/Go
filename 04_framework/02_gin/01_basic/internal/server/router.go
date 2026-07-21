package server

import (
	"ginlearn/internal/notes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(databasePool *pgxpool.Pool) *gin.Engine{
	r:=gin.Default()
	gin.SetMode(gin.ReleaseMode)

	if err := r.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	

	r.GET("/health",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"ok":true,
			"status":"healthy",

		})
	})

	notes.RegisterRoutes(r,databasePool)

	return r;
}