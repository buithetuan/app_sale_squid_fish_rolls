package handlers

import (
	"Backend/dto"
	"Backend/middeleware"
	"Backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminHandler struct {
	adminService *services.AdminService
}

// Admin control product
func NewAdminHandler(adminService *services.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

func (h *AdminHandler) AddProductHandler(c *gin.Context) {
	userID, err := middeleware.GetUserIDFromToken(c)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusUnauthorized, err.Error())
		return
	}

	if err := h.adminService.CheckAdmin(userID); err != nil {
		dto.SendResponse(c, nil, http.StatusForbidden, err.Error())
		return
	}

	var addProductRes dto.AddProductRes

	if err := c.ShouldBindJSON(&addProductRes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.adminService.AddProductService(addProductRes.ProductId, addProductRes.Name, addProductRes.Stock, addProductRes.Price)

	if err != nil {
		dto.SendResponse(c, nil, http.StatusInternalServerError, err.Error())
		return
	}
	dto.SendResponse(c, nil, http.StatusCreated, "Product added")

}

func (h *AdminHandler) RemoveProductHandler(c *gin.Context) {
	userID, err := middeleware.GetUserIDFromToken(c)

	if err != nil {
		dto.SendResponse(c, nil, http.StatusUnauthorized, err.Error())
		return
	}

	if err := h.adminService.CheckAdmin(userID); err != nil {
		dto.SendResponse(c, nil, http.StatusForbidden, err.Error())
		return
	}

	var removeProductRes dto.RemoveProductRes

	if err := c.ShouldBindJSON(&removeProductRes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.adminService.DeleteProductService(removeProductRes.ProductId)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusInternalServerError, err.Error())
		return
	}
	dto.SendResponse(c, nil, http.StatusOK, "Product removed")
}

func (h *AdminHandler) UpdateProductStockHandler(c *gin.Context) {
	userID, err := middeleware.GetUserIDFromToken(c)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusUnauthorized, err.Error())
		return
	}
	if err := h.adminService.CheckAdmin(userID); err != nil {
		dto.SendResponse(c, nil, http.StatusForbidden, err.Error())
		return
	}

	var updateProductRes dto.UpdateProductRes

	if err := c.ShouldBindJSON(&updateProductRes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.adminService.UpdateProductService(updateProductRes.ProductId, updateProductRes.Stock, updateProductRes.Price)

	if err != nil {
		dto.SendResponse(c, nil, http.StatusInternalServerError, err.Error())
		return
	}
	dto.SendResponse(c, nil, http.StatusOK, "Product updated")
}

// Admin control order
func (h *AdminHandler) UpdateStatusOrdersHandler(c *gin.Context) {
	userID, err := middeleware.GetUserIDFromToken(c)
	if err != nil {
		dto.SendResponse(c, nil, http.StatusUnauthorized, err.Error())
		return
	}
	if err := h.adminService.CheckAdmin(userID); err != nil {
		dto.SendResponse(c, nil, http.StatusForbidden, err.Error())
		return
	}

	var updateStatus dto.UpdateOrderStatusRes

	if err := c.ShouldBindJSON(&updateStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.adminService.UpdateStatusOrderService(updateStatus.OrderId, updateStatus.OrderStatus, updateStatus.PaymentStatus, updateStatus.ShipStatus)

	if err != nil {
		dto.SendResponse(c, nil, http.StatusInternalServerError, err.Error())
		return
	}

	dto.SendResponse(c, nil, http.StatusOK, "Product updated")
}
