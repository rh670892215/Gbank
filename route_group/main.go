package main

import (
	gbank2 "Gbank/template/gbank"
	"net/http"
)

func main() {
	engine := gbank2.NewEngine()
	engine.GET("/", func(c *gbank2.Context) {
		c.String(200, "hello")
	})

	engine.GET("/hello", func(c *gbank2.Context) {
		c.String(200, "hello,you are %s", c.QueryGet("name"))
	})

	engine.POST("/login", func(c *gbank2.Context) {
		c.JSON(200, gbank2.H{
			"name":   c.QueryPost("name"),
			"passwd": c.QueryPost("passwd"),
		})
	})

	engine.GET("/hello/:name", func(c *gbank2.Context) {
		c.String(200, "your name is %s", c.GetParam("name"))
	})

	engine.GET("/assets/*filepath", func(c *gbank2.Context) {
		c.JSON(http.StatusOK, gbank2.H{"filepath": c.GetParam("filepath")})
	})

	v1 := engine.NewGroup("/v1")
	v1.GET("/hello", func(c *gbank2.Context) {
		c.String(200, "v1 hello,you are %s", c.QueryGet("name"))
	})

	v2 := engine.NewGroup("/v2")
	v2.GET("/hello", func(c *gbank2.Context) {
		c.String(200, "v2 hello,you are %s", c.QueryGet("name"))
	})

	engine.Run(":9999")
}
