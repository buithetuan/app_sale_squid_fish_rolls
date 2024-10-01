package middeleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
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

func saveTokenToRedis(tokenKey string, token string, expiration time.Duration) error {
	err := redisClient.Set(ctx, tokenKey, token, expiration).Err() // Lưu token với khóa là tokenKey
	if err != nil {
		return fmt.Errorf("Error saving token to Redis: %v", err)
	}
	return nil
}

func CreateToken(data map[string]interface{}) (string, string, error) {
	userID := fmt.Sprintf("%v", data["user_id"])

	if err := redisClient.Del(ctx, userID+":access").Err(); err != nil {
		return "", "", fmt.Errorf("error deleting old access token from Redis: %v", err)
	}
	if err := redisClient.Del(ctx, userID+":refresh").Err(); err != nil {
		return "", "", fmt.Errorf("error deleting old refresh token from Redis: %v", err)
	}

	accessToken, err := createAccessToken(data)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := createRefreshToken(data)
	if err != nil {
		return "", "", err
	}

	if err := saveTokenToRedis(userID+":access", accessToken, 1*time.Hour); err != nil {
		log.Println("Error saving access token to Redis:", err)
	}
	if err := saveTokenToRedis(userID+":refresh", refreshToken, 7*24*time.Hour); err != nil {
		log.Println("Error saving refresh token to Redis:", err)
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

	userID := claims["user_id"]
	existingToken, err := redisClient.Get(ctx, fmt.Sprintf("%v:access", userID)).Result()
	if err != nil || existingToken == "" {
		return nil, fmt.Errorf("token does not exist in Redis")
	}

	return claims, nil
}

func GetUserIDFromToken(c *gin.Context) (uint, error) {
	token := c.GetHeader("Authorization")
	if token == "" {
		return 0, fmt.Errorf("No Authorization Token")
	}

	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}

	claims, err := DecodeToken(token)
	if err != nil {
		return 0, err
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("No User Token")
	}

	return uint(userIDFloat), nil
}
