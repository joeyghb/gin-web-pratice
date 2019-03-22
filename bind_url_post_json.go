package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
)

// 定义一个 Person 结构体，用来绑定数据
type Person struct {
	Name     string    `form:"name"`
	Message  string    `form:"message"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func bindPostData(c *gin.Context) {
	var person Person
	// 绑定到 person
	if c.ShouldBind(&person) != nil {
		c.JSON(http.StatusOK, gin.H{
			"Name":     person.Name,
			"Message":  person.Message,
			"Birthday": person.Birthday,
		})
	}
}

func main() {
	router := gin.Default()
	router.POST("/testing", bindPostData)
	router.Run(":18080")
}
