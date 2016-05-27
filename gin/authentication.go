package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func responseWithError(code int, message string, c *gin.Context) {
	resp := map[string] string{
		"error": message,
	}

	c.JSON(code, resp)
	c.Abort()
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("api_token")

		if token == "" {
			responseWithError(401, "API token required", c)
			return
		}

		if token != os.Getenv("API_TOKEN") {
			responseWithError(401, "Invalid API token", c)
			return
		}
		c.Next()
	}
}

func main() {
	os.Setenv("API_TOKEN", "token_value")

	r := gin.Default()
	r.Use(TokenAuthMiddleware())

	r.GET("/ranking", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"1" : "Mitsubishi",
			"2" : "Hitachi",
			"3" : "Sony",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"date" : "2016/05/27",
			"title" : "hogehoge",
			"detail" : "pepepepe",
		})
	})
	r.Run()
}
