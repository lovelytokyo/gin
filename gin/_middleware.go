package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"net/http"
)

func MyBenchLogger() {

}

func AuthRequired() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 認証済み
		// c.Next()

		// 未認証
		c.Redirect(http.StatusNonAuthoritativeInfo, "/login")
	}
}

/* Using middleware */
func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middlware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/benchmark", MyBenchLogger, benchEndpoint)

	// Authorization group
	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", func(c *gin.Context) {})
		authorized.POST("/submit", func(c *gin.Context) {})
		authorized.POST("/read", func(c *gin.Context) {})

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", func(c *gin.Context) {})

		r.Run(":8080")
	}
}