package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"os"
	"log"
	"io"
)

func main() {
	// Creates a gin router with default middleware;
	// logger and recovery middleware
	router := gin.Default();

	/*
	router.GET("/somdGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somPatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)
	*/

	/* Parameters in path */
	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)

	})

	/* Querystring parameters */
	router.GET("welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	/* Multipart/Urlencoded Form */
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})

	/* query + post form */
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		c.JSON(200, gin.H{
			"id": id,
			"page": page,
			"name": name,
			"message": message,
		})
	})

	/* put */
	router.PUT("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		name := c.PostForm("name")

		// なんらかの更新処理

		c.JSON(200, gin.H{
			"status": "putted",
			"id": id,
			"name": name,
		})
	})

	/* delete */
	router.DELETE("/user/:id", func(c *gin.Context) {
		id := c.Param("id")

		// なんらかの削除処理

		c.JSON(200, gin.H{
			"status": "deleted",
			"id": id,
		})
	})

	/* upload file */
	router.POST("/upload", func(c *gin.Context){
		file, header, err := c.Request.FormFile("upload")

		filename := header.Filename
		fmt.Println(header.Filename)
		out, err := os.Create("./tmp/"+filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
	})

	/* grouping routes */
	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			fmt.Println("router group 1 /login ")
		})
		v1.POST("/submit", func(c *gin.Context) {
			fmt.Println("router group 1 /submit ")
		})
		v1.POST("/read", func(c *gin.Context) {
			fmt.Println("router group 1 /read ")
		})
	}

	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			fmt.Println("router group 2 /login ")
		})
		v2.POST("/submit", func(c *gin.Context) {
			fmt.Println("router group 2 /submit ")
		})
		v2.POST("/read", func(c *gin.Context) {
			fmt.Println("router group 2 /read ")
		})
	}

	// By default it serves on :8080 unless a PORT environment variable was defined
	router.Run()
}
