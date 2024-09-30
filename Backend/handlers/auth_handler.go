package handlers

import (
	"Backend/dto"
	"Backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) SignUpHandler(c *gin.Context) {
	var signupRes dto.SignupRes

	if err := c.ShouldBind(&signupRes); err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.authService.SignUpService(signupRes.Username, signupRes.Password, signupRes.PhoneNumber, signupRes.Email, signupRes.Address); err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, err.Error())
		return
	}

	dto.SendResponse(c, nil, http.StatusOK, "User signed up successfully")
}

func (h *AuthHandler) LogInHandler(c *gin.Context) {
	var loginRes dto.LoginRes

	if err := c.ShouldBind(&loginRes); err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, "Invalid input")
		return
	}

	accessToken, refreshToken, err := h.authService.LoginService(loginRes.LoginType, loginRes.Username, loginRes.Password)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, err.Error())
		return
	}

	dto.SendResponse(c, dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken},
		http.StatusOK, "User logged in successfully")
}

func (h *AuthHandler) RefreshTokenHandler(c *gin.Context) {
	var tokenRes dto.RefreshTokenRes
	if err := c.ShouldBindJSON(&tokenRes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	accessToken, refreshToken, err := h.authService.RefreshTokenService(tokenRes.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	dto.SendResponse(c, dto.LoginResponse{accessToken, refreshToken}, http.StatusOK, "User refreshed successfully")
}
