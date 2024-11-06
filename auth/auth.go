package auth

import (
	"miniProject/models"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func HomeHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func LoginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
		"content": "",
	})
}

func LoginPostHandler(c *gin.Context){
	var credential models.Login
	err := c.Bind(&credential)
	if err != nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content": "Username or Password is invalid",
		})
	}

	if credential.Username != models.USER || credential.Password != models.PASSWORD {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"content": "Username or Password is invalid",
		})
	} else {
		//token

		claim := jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer: "book inventory",
			IssuedAt: time.Now().Unix(),
		}

		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		signedToken, err := sign.SignedString([]byte(models.SECRET))
		if err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"content": "Username or Password is invalid",
			})
			c.Abort()
		}

		q := url.Values{}
		q.Set("auth", signedToken)
		location := url.URL{
			Path:     "/books",
			RawQuery: q.Encode(),
		}

		c.Redirect(http.StatusMovedPermanently, location.RequestURI())
	}
}