package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func (c *gin.Context) {
		t := time.Now()

		// set example valiable
		c.Set("example", "12345")

		c.Next()

		// after request
		latency := time.Since(t)
		log.Println(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main()  {
	r := gin.New()
	r.Use(Logger()) // custom middlware

	r.GET("/test", func (c *gin.Context) {
		example := c.MustGet("example").(string)

		log.Println(example)
	})

	r.Run(":8080")
}
