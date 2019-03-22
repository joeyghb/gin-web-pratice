package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")

	router.GET("/main/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
			"message": "Hello Go Gin!!",
		})
	})

	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/posts.tmpl", gin.H{
			"title": "Posts",
		})
	})

	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/users.tmpl", gin.H{
			"title": "Users",
		})
	})

	router.Run(":18080")
}