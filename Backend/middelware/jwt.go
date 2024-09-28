package middeleware

import (
	"context"
	"fmt"
	"log"
	"time"

	"Backend/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
)

var (
	ctx         = context.Background()
	cfg         *config.Config
	redisClient *redis.Client
)

func init() {
	cfg = config.LoadConfig()

	redisClient = config.ConnectRedis(cfg)
}

func saveTokenToRedis(token string, userID string, expiration time.Duration) error {
	err := redisClient.Set(ctx, token, userID, expiration).Err()
	if err != nil {
		return fmt.Errorf("Error saving token to Redis: %v", err)
	}
	return nil
}

func CreateToken(data map[string]interface{}) (string, string, error) {
	accessToken, err := createAccessToken(data)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := createRefreshToken(data)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func createAccessToken(data map[string]interface{}) (string, error) {
	expire := time.Now().Add(1 * time.Hour)
	claims := jwt.MapClaims{
		"exp": expire.Unix(),
	}

	for key, value := range data {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(cfg.JWTSecretKey))
	if err != nil {
		return "", err
	}
	err = saveTokenToRedis(
		tokenStr,
		fmt.Sprintf("%v", data["user_id"]),
		1*time.Hour,
	)
	if err != nil {
		log.Println("Error saving access token to Redis:", err)
	}

	return tokenStr, nil
}

func createRefreshToken(data map[string]interface{}) (string, error) {
	expire := time.Now().Add(7 * 24 * time.Hour)
	claims := jwt.MapClaims{
		"exp": expire.Unix(),
	}

	for key, value := range data {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(cfg.JWTSecretKey))
	if err != nil {
		return "", err
	}

	err = saveTokenToRedis(
		tokenStr,
		fmt.Sprintf("%v", data["user_id"]),
		7*24*time.Hour,
	)
	if err != nil {
		log.Println("Error saving refresh token to Redis:", err)
	}
	return tokenStr, nil
}

func DecodeToken(tokenString string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWTSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return claims, nil
}
