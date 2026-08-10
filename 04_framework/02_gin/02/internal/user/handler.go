package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandle(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var userCreationRequest UserCreationRequest;
	err:=c.ShouldBindJSON(&userCreationRequest);
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"Please Provide Proper email and password in proper json format",
			"error":err.Error(),
		})
		return;
	}

	if userCreationRequest.Email=="" || userCreationRequest.Password==""{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"Email or Password missing",
		})
		return;
	}

	res, err := h.service.CreateUser(c.Request.Context(), userCreationRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, res)


}

func (h *Handler) Login(c *gin.Context){
    var loginReqObj UserLoginRequest
	err:=c.ShouldBindJSON(&loginReqObj)

	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"Please Provide Proper email and password in proper json format",
			"error":err.Error(),
		})
		return;
	}

	loginResObj,err:=h.service.Login(c.Request.Context(),loginReqObj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted,gin.H{
		"data":loginResObj,
	})

}