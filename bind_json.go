package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Define Req Json struct , for binding Req Data
type Req struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Define Rsp Json struct , for binding Rsp Data
type Rsp struct {
	Name     string
	Message  string
}

//Post {"user": "Pgluffy","password": "123"}
func bindData(c *gin.Context) {
	var reqjson Req
	var rspjson Rsp

	if  err := c.ShouldBindJSON(&reqjson); err == nil {

		if reqjson.User == "Pgluffy" && reqjson.Password == "123" {

			rspjson.Name = reqjson.User
			rspjson.Message = "you are logged in"

			c.JSON(http.StatusOK, gin.H{
				"Name":     rspjson.Name,
				"Message":  rspjson.Message,
			})
		}else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Get Req Data Error"})
	}
}

func main() {
	router := gin.Default()
	router.POST("/testbind", bindData)
	router.Run(":18080")
}
