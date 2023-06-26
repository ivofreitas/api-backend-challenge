package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sword/api-backend-challenge/config"
	"github.com/sword/api-backend-challenge/model"
	"strings"
)

func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasSuffix(c.Path(), "/health") || strings.HasSuffix(c.Path(), "/login") || strings.HasSuffix(c.Path(), "/user") {
			return next(c)
		}

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			responseErr := model.ErrorDiscover(model.Unauthorized{DeveloperMessage: "Missing authorization header"})
			return c.JSON(responseErr.StatusCode, responseErr)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			responseErr := model.ErrorDiscover(model.Unauthorized{DeveloperMessage: "Invalid token format"})
			return c.JSON(responseErr.StatusCode, responseErr)
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetEnv().Authorization.Secret), nil
		})

		if err != nil || !token.Valid {
			responseErr := model.ErrorDiscover(model.Unauthorized{DeveloperMessage: "Make sure the header parameter Authorization is valid"})
			return c.JSON(responseErr.StatusCode, responseErr)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			responseErr := model.ErrorDiscover(model.Unauthorized{DeveloperMessage: "Missing JWT Claims"})
			return c.JSON(responseErr.StatusCode, responseErr)
		}

		username, ok := claims["user"].(string)
		if !ok {
			responseErr := model.ErrorDiscover(model.Unauthorized{DeveloperMessage: "Missing JWT Username"})
			return c.JSON(responseErr.StatusCode, responseErr)
		}

		role, ok := claims["role"].(string)
		if !ok {
			responseErr := model.ErrorDiscover(model.Unauthorized{DeveloperMessage: "Missing JWT Role"})
			return c.JSON(responseErr.StatusCode, responseErr)
		}

		c.Set("user", username)
		c.Set("role", role)

		return next(c)
	}
}

func CheckRole(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role := c.Get("role").(string)
			role = strings.ToLower(role)

			isAllowed := false
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					isAllowed = true
					break
				}
			}

			if !isAllowed {
				responseErr := model.ErrorDiscover(model.Forbidden{})
				return c.JSON(responseErr.StatusCode, responseErr)
			}

			return next(c)
		}
	}
}
