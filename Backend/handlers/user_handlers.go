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

func (h *UserHandler) UserAddItemsHandler(c *gin.Context) {
	userID, err := middeleware.GetUserIDFromToken(c)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusUnauthorized, err.Error())
		return
	}

	var addItemsRes dto.AddItemRes
	if err := c.ShouldBind(&addItemsRes); err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.userService.UserAddItemsToCart(c.Request.Context(), userID, addItemsRes.ProductId, addItemsRes.Quantity); err != nil {
		dto.SendResponse(c, nil, http.StatusInternalServerError, err.Error())
		return
	}
	dto.SendResponse(c, nil, http.StatusOK, "Added items successfully")
}

func (h *UserHandler) UserRemoveItemsHandler(c *gin.Context) {
	userID, err := middeleware.GetUserIDFromToken(c)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusUnauthorized, err.Error())
	}

	var removeItemsRes dto.RemoveItemRes
	if err := c.ShouldBind(&removeItemsRes); err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, "Invalid input")
	}

	if err := h.userService.UserRemoveItemFromCart(c.Request.Context(), userID, removeItemsRes.ProductId); err != nil {
		dto.SendResponse(c, nil, http.StatusInternalServerError, err.Error())
		return
	}
	dto.SendResponse(c, nil, http.StatusOK, "Removed items successfully")
}

func (h *UserHandler) UserOrderHandler(c *gin.Context) {
	userID, err := middeleware.GetUserIDFromToken(c)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusUnauthorized, err.Error())
	}

	var orderRes dto.CreateOrderRes
	if err := c.ShouldBind(&orderRes); err != nil {
		dto.SendResponse(c, nil, http.StatusBadRequest, "Invalid input")
		return
	}
	orderResponse, err := h.userService.UserOrderService(c, userID, orderRes.PaymentMethod, orderRes.DeliveryAddress)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusInternalServerError, err.Error())
		return
	}
	dto.SendResponse(c, orderResponse, http.StatusOK, "Order successfully")
}
