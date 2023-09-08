package main

import (
	"Gbank/template/gbank"
	"fmt"
	"html/template"
	"time"
)

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	engine := gbank.NewEngine()
	engine.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	engine.LoadHTMLTemplate("templates/*")
	engine.Static("/assets", "./static")

	engine.GET("/", func(c *gbank.Context) {
		c.HTML(200, "css.tmpl", nil)
	})

	engine.Run(":9999")
}
