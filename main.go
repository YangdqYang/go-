package main

import (
	 "github.com/gin-gonic/gin"
	 "fmt"
	 "net/http"
) 

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main()  {
	r := gin.Default()
	r.GET("/string", func(c *gin.Context) {
		c.String(http.StatusOK, "it works")
	})

	r.GET("/json", func(c *gin.Context) {
		data := Person{Name: "yang", Age: 12}
		c.JSON(http.StatusOK, gin.H{
			"errorCode": 0,
			"data" : data,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
	fmt.Println("hello world")
}