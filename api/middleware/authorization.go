package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sword/api-backend-challenge/config"
	"github.com/sword/api-backend-challenge/model"
	"net/http"
)

func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Path() == "/health" || c.Path() == "/login" {
			return next(c)
		}

		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return model.ErrorDiscover(model.Unauthorized{})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetEnv().Authorization.Secret), nil
		})

		if err != nil || !token.Valid {
			return model.ErrorDiscover(model.Unauthorized{})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return model.ErrorDiscover(model.Unauthorized{})
		}

		username, ok := claims["user"].(string)
		if !ok {
			return model.ErrorDiscover(model.Unauthorized{})
		}

		role, ok := claims["role"].(string)
		if !ok {
			return model.ErrorDiscover(model.Unauthorized{})
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

			// Check if the user's role is allowed to access the endpoint
			isAllowed := false
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					isAllowed = true
					break
				}
			}

			if !isAllowed {
				return c.String(http.StatusForbidden, "Access denied")
			}

			return next(c)
		}
	}
}
