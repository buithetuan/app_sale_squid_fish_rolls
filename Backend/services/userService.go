package services

import (
	"Backend/middelware"
	"Backend/models"
	"Backend/repo"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService(userRepo *repo.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) SignUpService(username string, password string) error {
	existingUser, err := s.userRepo.GetUserByUserName(username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return fmt.Errorf("username %s already exists", username)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	user := models.Users{Username: username, Password: string(hash)}

	if err := s.userRepo.CreateUser(&user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) LoginService(loginType string, username string, password string) (string, string, error) {
	switch loginType {
	//case "google":
	//	return s.LoginWithGoogle(token)
	case "username":
		return s.LoginWithUsername(username, password)

	default:
		return "", "", errors.New("unsupported loginType")
	}
}

//func (s *UserService) LoginWithGoogle(token) error {
//}

func (s *UserService) LoginWithUsername(username string, password string) (string, string, error) {
	existingUser, err := s.userRepo.GetUserByUserName(username)
	if err != nil {
		return "", "", err
	}

	if existingUser == nil {

		return "", "", fmt.Errorf("username %s does not exists", username)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password)); err != nil {
		return "", "", fmt.Errorf("invalid password")
	}

	data := map[string]interface{}{
		"user_id": existingUser.UserID,
	}

	accessToken, refreshToken, err := middeleware.CreateToken(data)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (s *UserService) UpdateUserService(userID int, email string, phoneNumber string, address string) (*models.Users, error) {
	existingUse, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	if existingUse == nil {
		return nil, fmt.Errorf("user %d does not exists", userID)
	}

	existingUse.Email = email
	existingUse.PhoneNumber = phoneNumber
	existingUse.Address = address

	updatedUser, err := s.userRepo.UpdateUser(userID, existingUse)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *UserService) GetUserService(userID int) (*models.Users, error) {
	return s.userRepo.GetUserByID(userID)
}
