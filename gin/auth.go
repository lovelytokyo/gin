package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var secrets = gin.H{
	"foo": gin.H{
		"email": "foo@bar.com",
		"phone": "123433",
	},
	"austin": gin.H{
		"email": "austin@bar.com",
		"phone": "66666",
	},
}

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "token required"})
			c.Abort()
		}

		os.Setenv("TOKEN", "tokentest")
		if token != os.Getenv("TOKEN") {
			c.Redirect(http.StatusMovedPermanently, "/readme")
			c.Abort()
		}
		c.Next()
	}
}

func main () {
	r := gin.Default()

	// basic auth
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo": "bar",
		"austin": "1234",
		"lena": "hello2",
		"manu": "4321",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// custom auth
	designer := r.Group("/designer", TokenAuth())
	designer.GET("/news", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"news": "hogehoge"})
	})

	// allow no auth
	r.GET("/readme", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"readme": "thanks"})
	})

	r.Run(":8080")
}
