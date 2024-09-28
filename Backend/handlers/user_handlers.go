package handlers

import (
	"Backend/dto"
	"Backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) SignUpHandler(c *gin.Context) {
	var signupRes dto.SignupRes

	if err := c.ShouldBind(&signupRes); err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.userService.SignUpService(signupRes.Username, signupRes.Password); err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, err.Error())
		return
	}

	dto.SendResponse(c, nil, http.StatusOK, "User signed up successfully")
}

func (h *UserHandler) LogInHandler(c *gin.Context) {
	var loginRes dto.LoginRes

	if err := c.ShouldBind(&loginRes); err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, "Invalid input")
		return
	}

	accessToken, refreshToken, err := h.userService.LoginService(loginRes.LoginType, loginRes.Username, loginRes.Password)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, err.Error())
		return
	}

	dto.SendResponse(c, dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken},
		http.StatusOK, "User logged in successfully")
}

func (h *UserHandler) UpdateUserHandler(c *gin.Context) {
	var updateUser dto.UpdateUserRes

	if err := c.ShouldBind(&updateUser); err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, "Invalid input")
		return
	}

	userInfo, err := h.userService.UpdateUserService(updateUser.UserID, updateUser.Email, updateUser.PhoneNumber, updateUser.Address)

	if err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, err.Error())
		return
	}

	dto.SendResponse(c, dto.UserInfoResponse{
		Username:    userInfo.Username,
		Email:       userInfo.Email,
		PhoneNumber: userInfo.PhoneNumber,
		Address:     userInfo.Address,
	}, http.StatusOK, "User updated successfully")
}

func (h *UserHandler) GetUserDetailHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, "Invalid user ID")
		return
	}

	userInfo, err := h.userService.GetUserService(userID)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusNotFound, "User not found")
		return
	}

	dto.SendResponse(c, dto.UserInfoResponse{
		Username:    userInfo.Username,
		Email:       userInfo.Email,
		PhoneNumber: userInfo.PhoneNumber,
		Address:     userInfo.Address,
	}, http.StatusOK, "User details retrieved successfully")
}
