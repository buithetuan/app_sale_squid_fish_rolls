package middeleware

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/go-redis/redis/v8"
    "Backend/config"
)

var (
    cfg          *config.Config
    redisClient  *redis.Client
)

func init() {
    cfg = config.LoadConfig()

    redisClient = config.ConnectRedis(cfg)
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
    return token.SignedString([]byte(cfg.JWTSecretKey))
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
    return token.SignedString([]byte(cfg.JWTSecretKey))
}
