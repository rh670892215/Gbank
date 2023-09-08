package main

import (
	"Gbank/middleware/gbank"
	"net/http"
)

func main() {
	r := gbank.Default()
	r.GET("/", func(c *gbank.Context) {
		c.String(http.StatusOK, "Hello bank\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gbank.Context) {
		names := []string{"bank"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
