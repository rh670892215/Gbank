package main

import (
	"Gbank/route_group/gbank"
	"net/http"
)

func main() {
	engine := gbank.NewEngine()
	engine.GET("/", func(c *gbank.Context) {
		c.String(200, "hello")
	})

	engine.GET("/hello", func(c *gbank.Context) {
		c.String(200, "hello,you are %s", c.QueryGet("name"))
	})

	engine.POST("/login", func(c *gbank.Context) {
		c.JSON(200, gbank.H{
			"name":   c.QueryPost("name"),
			"passwd": c.QueryPost("passwd"),
		})
	})

	engine.GET("/hello/:name", func(c *gbank.Context) {
		c.String(200, "your name is %s", c.GetParam("name"))
	})

	engine.GET("/assets/*filepath", func(c *gbank.Context) {
		c.JSON(http.StatusOK, gbank.H{"filepath": c.GetParam("filepath")})
	})

	v1 := engine.NewGroup("/v1")
	v1.GET("/hello", func(c *gbank.Context) {
		c.String(200, "v1 hello,you are %s", c.QueryGet("name"))
	})

	v2 := engine.NewGroup("/v2")
	v2.GET("/hello", func(c *gbank.Context) {
		c.String(200, "v2 hello,you are %s", c.QueryGet("name"))
	})

	engine.Run(":9999")
}
