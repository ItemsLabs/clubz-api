package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const (
	UserIDKey = "user_id"
)

// userID retrieves the user ID from the context, set by the middleware.
func userID(c echo.Context) string {
	userID, ok := c.Get(UserIDKey).(string)
	if !ok || userID == "" {
		log.Println("Failed to retrieve user_id from context or user_id is empty")
		return ""
	}
	return userID
}

// JWTMiddleware checks for a valid JWT token in the request header.

func (e *Env) JWTMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
			if tokenString == "" {
				log.Println("JWT Middleware: Token string is empty")
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
			}

			token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil {
				log.Printf("JWT Middleware: Error parsing token - %v", err)
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
				userID, ok := (*claims)["user_id"].(string)
				if !ok {
					log.Println("JWT Middleware: user_id claim is missing or not a string")
					return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
				}
				// check if the user actually exists in the db
				if _, err := e.Store.GetUserByID(userID); err != nil {
					log.Printf("JWT Middleware: Error retrieving user - %v", err)
					return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
				}
				c.Set("user_id", userID)
			} else {
				log.Println("JWT Middleware: Token is not valid")
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			return next(c)
		}
	}
}
