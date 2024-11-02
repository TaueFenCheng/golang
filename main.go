package main

import (
	"fmt"
	"net/http"
	"strconv"

	SubModule "demo/submodule"

	"github.com/gin-gonic/gin"

	"main2"
)

func Demo() {
	fmt.Println("from demo module")
}

func main() {
	fmt.Println("world")

	// 引用同一个项目中的不同文件夹下的包
	SubModule.SubModule()
	SubModule.SubModule2()
	// 不同go 文件下 方法

	SubModule.Sub2Module()

	// 引用另一个项目中的包
	main2.Main2Module()

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
