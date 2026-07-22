package notes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(r *gin.Engine,pool *pgxpool.Pool){
	repo:=NewPostgresRepository(pool)
	service:=NewService(repo)
	handler:=NewHandler(service)

	notesGroup:=r.Group("/notes")
	{
		notesGroup.POST("",handler.Create)
		notesGroup.GET("",handler.List)
	}
	
}