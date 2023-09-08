package main

import "Gbank/base/gbank"

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

	engine.Run(":9999")
}
