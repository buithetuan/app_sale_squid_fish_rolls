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
func (r *UserRepo) GetUserByID(userID uint) (*models.Users, error) {
	var user models.Users
	if err := r.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return &user, nil
}

func (r *UserRepo) UpdateUser(userID uint, userUpdates *models.Users) (*models.Users, error) {
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

func (r *UserRepo) IsAdmin(userID uint) (bool, error) {
	var user models.Users
	if err := r.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("failed to get user by id: %v", err)
	}

	return user.IsAdmin, nil
}
func (r *UserRepo) AddToTotalBuy(userID uint, finalPrice float64) error {
	if err := r.db.Model(&models.Users{}).Where("user_id = ?", userID).Update("total_buy", gorm.Expr("total_buy + ?", finalPrice)).Error; err != nil {
		return fmt.Errorf("failed to update total buy for user %d: %v", userID, err)
	}
	return nil
}

func (r *UserRepo) UpdateUserRank(userID uint) error {
	user, err := r.GetUserByID(userID)
	if err != nil {
		return err
	}

	var rank string

	switch {
	case user.TotalBuy >= 3000000.0:
		rank = "silver"
	case user.TotalBuy >= 5000000.0:
		rank = "gold"
	case user.TotalBuy >= 10000000.0:
		rank = "premium"
	case user.TotalBuy >= 15000000.0:
		rank = "patron"
	default:
		rank = "bronze"
	}

	if err := r.db.Model(&models.Users{}).Where("user_id = ?", userID).Update("rank", rank).Error; err != nil {
		return fmt.Errorf("failed to update user rank for user %d: %v", userID, err)
	}

	return nil
}
func (r *UserRepo) GetRankFromUserID(userID uint) (models.UserRank, error) {
	var user models.Users
	if err := r.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return "", err
	}
	return user.Rank, nil
}
