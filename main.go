package main

import (
	 "github.com/gin-gonic/gin"
	 "fmt"
	 "net/http"
) 

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Dogs []Dog `json:"dogs"`
}

type Dog struct {
	Name string `json:"name"`
	Cag string `json:"cag"`
}

func main()  {
	Hello()

	r := gin.Default()
	r.GET("/user/:name", func(c *gin.Context) {
		type user struct {
			Name string `json:"name"`
			Age string `json:"age"`
			Height string `json:"height"`
		}
		name := c.Param("name")
		age := c.Query("age")
		height := c.Query("height")

		data := user{Name: name, Age: age, Height: height}
		c.JSON(http.StatusOK, gin.H{
			"errorCode": 0,
			"data" : data,
		})
	})



	r.GET("/string", func(c *gin.Context) {
		c.String(http.StatusOK, "it works")
	})

	r.GET("/json", func(c *gin.Context) {
		var s []Dog
		s = append(s, Dog{Name: "小贝贝", Cag: "泰迪"})
		s = append(s, Dog{Name: "小嘿嘿", Cag: "泰迪"})
		s = append(s, Dog{Name: "白丹", Cag: "牧羊犬"})

		data := Person{Name: "yang", Age: 12, Dogs: s}

		c.JSON(http.StatusOK, gin.H{
			"errorCode": 0,
			"data" : data,
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
	fmt.Println("hello world")
}