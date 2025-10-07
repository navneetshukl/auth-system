package middleware

import (
	"auth-system/internal/core/user"
	ratelimiting "auth-system/pkg/rateLimit"
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

var rateLimiter = ratelimiting.NewTokenBucket(1, 5)

// RateLimitMiddleware limits requests per user/email
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdentifier := c.GetString("username")

		if userIdentifier == "" {
			userIdentifier = c.ClientIP()
		}

		if !rateLimiter.AllowRequest(userIdentifier) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":  "too many requests, please try again later",
				"status": http.StatusTooManyRequests,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}