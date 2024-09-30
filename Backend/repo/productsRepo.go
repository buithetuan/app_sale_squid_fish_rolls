package repo

import (
	"Backend/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) GetProductByID(productID string) (*models.Products, error) {
	var product models.Products
	if err := r.db.First(&product, "product_id = ?", productID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("fail to get product by id : %w", err)
	}
	return &product, nil
}

func (r *ProductRepo) AddProduct(product *models.Products) error {
	if err := r.db.Create(product).Error; err != nil {
		return fmt.Errorf("ProductRepo AddProduct err: %v", err)
	}
	return nil
}

func (r *ProductRepo) DeleteProduct(productID string) error {
	if err := r.db.Where("id = ?", productID).Delete(&models.Products{}).Error; err != nil {
		return fmt.Errorf("ProductRepo DeleteProduct err: %v", err)
	}
	return nil
}

func (r *ProductRepo) UpdateProduct(productID string, productUpdates *models.Products) (*models.Products, error) {
	var product models.Products

	if err := r.db.Model(product).Where("product_id = ?", productID).Select("stock", "price").Updates(productUpdates).Error; err != nil {
		return nil, fmt.Errorf("fail to update product: %v", err)
	}

	updateProduct, err := r.GetProductByID(productID)
	if err != nil {
		return nil, fmt.Errorf("fail to get product: %v", err)
	}
	if updateProduct == nil {
		return nil, fmt.Errorf("product does not exist")
	}

	return updateProduct, nil
}
