package notes

import (
	"net/http"
	"strconv"
	// "strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Create(c *gin.Context) {

	var req CreateNoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	note, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, note)
}

func (h *Handler)List(c *gin.Context){
	notes,err:=h.service.List(c.Request.Context())
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return;
	}

	c.JSON(http.StatusOK,gin.H{
		"list":notes,
	})
}

func (h *Handler) GetNoteById(c *gin.Context){
	id:=c.Query("id");
	iId,err:=strconv.Atoi(id);
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
	}

	note,err:=h.service.GetNoteById(c.Request.Context(),iId);
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return;
	}

	c.JSON(http.StatusOK,gin.H{
		"status":true,
		"data":note,
	})
}

/*
func (h *Handler) GetByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	note, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "note not found",
		})
		return
	}

	c.JSON(http.StatusOK, note)
}
*/