package repo

import (
	"Backend/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user *models.Users) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func (r *UserRepo) GetUserByUserName(username string) (*models.Users, error) {
	var user models.Users

	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	return &user, nil
}
func (r *UserRepo) GetUserByID(userID int) (*models.Users, error) {
	var user models.Users
	if err := r.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return &user, nil
}

func (r *UserRepo) UpdateUser(userID int, userUpdates *models.Users) (*models.Users, error) {
	var user models.Users

	if err := r.db.Model(user).Where("user_id = ?", userID).Select("email", "phone_number", "address").Updates(userUpdates).Error; err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	updatedUser, err := r.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %v", err)
	}

	if updatedUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	return updatedUser, nil
}
