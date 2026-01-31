package token

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go_fiber_wibowo/app/module/auth/response"
	"go_fiber_wibowo/utils/config"
	"strings"
	"time"
)

type UserClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type ResponseToken struct {
	UserId  string `json:"user_id"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ExtractToken(c *fiber.Ctx) string {
	bearerToken := c.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(tokenString string) (*response.LoginResponse, error) {
	Cfg := config.NewConfig()
	dataToken := &response.LoginResponse{}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Cfg.Middleware.Jwt.Secret), nil
	})

	if err != nil {
		fmt.Println("err != nil :", token)
		return dataToken, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		dataToken.UserId = fmt.Sprintf("%v", claims["user_id"])
		dataToken.Token = tokenString

		return dataToken, nil
	}

	return dataToken, nil
}

func GenerateToken(userId uuid.UUID) (string, error) {
	Conf := config.NewConfig()

	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userId,
		"exp":        time.Now().Add(time.Hour * Conf.Middleware.Jwt.Expiration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(Conf.Middleware.Jwt.Secret))
}

func TokenValid(c *fiber.Ctx) (*jwt.Token, error) {
	Cfg := config.NewConfig()
	tokenString := ExtractToken(c)
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Cfg.Middleware.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return t, nil
}
