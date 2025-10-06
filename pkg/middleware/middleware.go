package middleware

import (
	"auth-system/internal/core/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("authToken")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":  "authentication token missing",
				"status": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &user.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return user.JWTSecret, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":  "invalid or expired token",
				"status": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*user.Claims); ok && token.Valid {
			// 4️⃣ Save the user info in the Gin context
			c.Set("username", claims.Username)
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error":  "unauthorized access",
			"status": http.StatusUnauthorized,
		})
		c.Abort()
	}
}
