package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User		string `form:"user" json:"user" binding:"required"`
	Password	string `form:"password" json: "password" binding:"required"`
}

func main()  {
	router := gin.Default()

	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if c.BindJSON(&json) == nil {
			if json.User == "manu" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			}else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		if c.Bind(&form) == nil {
			if form.User == "manu" && form.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			}else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	router.Run()
}
