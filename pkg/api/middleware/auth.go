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
	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Next()
}

/*
validateToken is for decrypting a jwt token using HMAC256 algorithm

Parameters:
- c: Gin Context.
*/
func validateToken(token string) error {
	fmt.Println("Token validating.........")
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		secret := viper.GetString("KEY")
		//secret := os.Getenv("KEY")
		return []byte(secret), nil
	})

	return err
}
