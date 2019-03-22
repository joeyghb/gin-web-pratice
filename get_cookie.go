package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

func getcookie(c *gin.Context) {

        cookie, err := c.Cookie("gin_cookie")

        if err != nil {
            cookie = "NotSet"
            c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
        }

        c.JSON(http.StatusOK, gin.H{"cookie": cookie})
		fmt.Printf("Cookie value: %s \n", cookie)
}


func main() {
	router := gin.Default()
	router.GET("/cookie", getcookie)
	router.Run(":18080")
}