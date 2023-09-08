package main

import (
	"Gbank/route_tire/gbank"
	"net/http"
)

func main() {
	engine := gbank.NewEngine()
	engine.Get("/", func(c *gbank.Context) {
		c.String(200, "hello")
	})

	engine.Get("/hello", func(c *gbank.Context) {
		c.String(200, "hello,you are %s", c.QueryGet("name"))
	})

	engine.POST("/login", func(c *gbank.Context) {
		c.JSON(200, gbank.H{
			"name":   c.QueryPost("name"),
			"passwd": c.QueryPost("passwd"),
		})
	})

	engine.Get("/hello/:name", func(c *gbank.Context) {
		c.String(200, "your name is %s", c.GetParam("name"))
	})

	engine.Get("/assets/*filepath", func(c *gbank.Context) {
		c.JSON(http.StatusOK, gbank.H{"filepath": c.GetParam("filepath")})
	})

	engine.Run(":9999")
}
