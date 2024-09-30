package handlers

import (
	"Backend/dto"
	"Backend/middeleware"
	"Backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) UpdateUserHandler(c *gin.Context) {
	userID, err := middeleware.GetUserIDFromToken(c)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusUnauthorized, err.Error())
		return
	}

	var updateUser dto.UpdateUserRes
	if err := c.ShouldBind(&updateUser); err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, "Invalid input")
		return
	}

	userInfo, err := h.userService.UpdateUserService(userID, updateUser.Email, updateUser.PhoneNumber, updateUser.Address)
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
	userID, err := middeleware.GetUserIDFromToken(c)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusUnauthorized, err.Error())
		return
	}

	userInfo, err := h.userService.GetUserService(userID)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusInternalServerError, err.Error())
		return
	}

	if userInfo == nil {
		dto.SendResponse(c, nil, http.StatusUnauthorized, "User not found")
		return
	}

	dto.SendResponse(c, dto.UserInfoResponse{
		Username:    userInfo.Username,
		Email:       userInfo.Email,
		PhoneNumber: userInfo.PhoneNumber,
		Address:     userInfo.Address,
	}, http.StatusOK, "User details retrieved successfully")
}

func (h *UserHandler) UserChosePaymentMethodHandler(c *gin.Context) {
	userID, err := middeleware.GetUserIDFromToken(c)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusUnauthorized, err.Error())
		return
	}
	var chosePaymentMethod dto.ChosePaymentMethodRes
	if err := c.ShouldBindJSON(&chosePaymentMethod); err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, "Invalid input")
		return
	}

	err = h.userService.UserChoseMethodPayService(userID, chosePaymentMethod.OrderId, chosePaymentMethod.Method)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusInternalServerError, err.Error())
		return
	}
	dto.SendResponse(c, nil, http.StatusOK, "Chose payment method successfully")
}
