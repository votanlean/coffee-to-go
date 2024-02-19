package main

import (
	"coffeeshop/coffee"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/home", func(c *gin.Context) {
		coffeelist, _ := coffee.GetCoffees()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"list": coffeelist.List,
		})
	})
	r.GET("/ping", ping)
	r.GET("/coffee", getCoffee)
	r.Run(":8080")
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getCoffee(c *gin.Context) {
	coffeelist, _ := coffee.GetCoffees()
	c.String(http.StatusOK, " %s", coffeelist)
}