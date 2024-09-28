package dto

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserInfoResponse struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

func SendResponse(c *gin.Context, data interface{}, status int, message string) {
	response := APIResponse{data, status, message}
	c.JSON(status, response)
}
