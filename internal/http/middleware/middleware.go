package middleware

import (
	"Trello/internal/auth/jwtutil"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func InfoLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestTime := time.Now()
		method := c.Request().Method
		path := c.Request().URL.Path
		err := next(c)
		responseTime := time.Since(requestTime)
		log.Printf("- %s request to %s took %v\n", method, path, responseTime)
		return err
	}
}

func IsAuthenticatedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		claims := &jwtutil.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtutil.JwtKey, nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"Error": "Invalid or expired token",
			})
		}

		c.Set("username", claims.Username)

		return next(c)
	}
}
