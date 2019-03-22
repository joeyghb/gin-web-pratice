package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		firstname := c.DefaultQuery("firstname", "None")
		lastname  := c.Query("lastname")

		c.JSON(http.StatusOK, gin.H{
			"name":      name,
			"action":    action,
			"firstname": firstname,
			"lastname":  lastname,
		})
	})
	r.Run(":18080")
}