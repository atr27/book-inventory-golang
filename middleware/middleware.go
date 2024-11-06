package middleware

import (
	"miniProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"fmt"
)

func AuthValidation(c *gin.Context)  {
	var tokenString string
	tokenString = c.Query("auth")
	if tokenString == "" {
		tokenString = c.PostForm("auth")
		if tokenString == "" {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"content": "Unauthorized",
			})
			c.Abort()
		}
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, invalid := token.Method.(*jwt.SigningMethodHMAC); !invalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}	
		return []byte(models.SECRET), nil
	})	
	
	if token != nil && err == nil {
		fmt.Println("Token Valid")
	} else {
		c.Abort()
	}
}