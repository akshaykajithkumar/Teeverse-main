package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

/*
AdminAuthMiddleware is a middleware for admin authentication

Parameters:
- c: Gin Context.
*/
func AdminAuthMiddleware(c *gin.Context) {
	token, _ := c.Cookie("Authorization")
	fmt.Println("Token::", token)
	fmt.Println(token)
	jwttoken, err := validateToken(token)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if err != nil || !jwttoken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		c.Abort()
		return
	}

	claims, ok := jwttoken.Claims.(jwt.MapClaims)
	if !ok || !jwttoken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		c.Abort()
		return
	}

	role, ok := claims["role"].(string)
	if !ok || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		c.Abort()
		return
	}
	c.Next()
}

/*
validateToken is for decrypting a jwt token using HMAC256 algorithm

Parameters:
- c: Gin Context.
*/
func validateToken(token string) (*jwt.Token, error) {
	fmt.Println("Token validating.........")
	jwttoken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		secret := viper.GetString("KEY")
		//secret := os.Getenv("KEY")
		return []byte(secret), nil
	})

	return jwttoken, err
}
