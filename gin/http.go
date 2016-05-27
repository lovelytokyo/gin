package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()

	s := &http.Server {
		Addr: ":3000",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 <<20,
	}
	router.GET("/", func (c *gin.Context) {
		time.Sleep(2000)

		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	s.ListenAndServe()

}
