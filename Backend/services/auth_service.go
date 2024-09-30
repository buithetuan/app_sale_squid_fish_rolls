package services

import (
	"Backend/middeleware"
	"Backend/models"
	"Backend/repo"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repo.UserRepo
	cartRepo *repo.CartRepo
}

func NewAuthService(userRepo *repo.UserRepo, cartRepo *repo.CartRepo) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		cartRepo: cartRepo,
	}
}

func (s *AuthService) SignUpService(username string, password string, email string, phone_number string, address string) error {
	existingUser, err := s.userRepo.GetUserByUserName(username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return fmt.Errorf("username %s already exists", username)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.Users{Username: username, Password: string(hash), Email: email, PhoneNumber: phone_number, Address: address}
	if err := s.userRepo.CreateUser(&user); err != nil {
		return err
	}

	cart := models.Carts{UserID: int(user.UserID)}
	if err := s.cartRepo.CreateCart(&cart); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) LoginService(loginType string, username string, password string) (string, string, error) {
	switch loginType {
	//case "google":
	//	return s.LoginWithGoogle(token)
	case "username":
		return s.LoginWithUsername(username, password)

	default:
		return "", "", errors.New("unsupported loginType")
	}
}

func (s *AuthService) LoginWithUsername(username string, password string) (string, string, error) {
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

func (s *AuthService) RefreshTokenService(refreshToken string) (string, string, error) {
	claims, err := middeleware.DecodeToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err := middeleware.CreateToken(claims)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
