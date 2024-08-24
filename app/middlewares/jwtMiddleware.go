package middlewares

import (
	"star-pos/app/configs"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(configs.JWT_SECRET),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.JWT_SECRET))
}

func ExtractTokenUserId(e echo.Context) string {
	header := e.Request().Header.Get("authorization")
	headerToken := strings.Split(header, " ")
	token := headerToken[len(headerToken)-1]
	tokenJWT, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.JWT_SECRET), nil
	})

	if tokenJWT.Valid {
		claims := tokenJWT.Claims.(jwt.MapClaims)
		userId, isValidUserId := claims["userId"].(string)
		if !isValidUserId {
			return ""
		}
		return userId
	}
	return ""
}
