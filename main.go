package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("world")
	i, err := strconv.Atoi("-42")
	s := strconv.Itoa(-42)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(i)
	fmt.Println(s)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
