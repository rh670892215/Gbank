package gbank

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type H map[string]interface{}

type Context struct {
	writer     http.ResponseWriter
	Req        *http.Request
	StatusCode int
	method     string
	path       string
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		writer: w,
		Req:    r,
		method: r.Method,
		path:   r.URL.Path,
	}
}

func (c *Context) QueryGet(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) QueryPost(key string) string {
	return c.Req.PostForm.Get(key)
}

func (c *Context) String(code int, format string, params ...string) {
	c.setCode(code)
	c.writer.Write([]byte(fmt.Sprintf(format, params)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.setHeader("Content-Type", "application/json")
	c.setCode(code)

	encoder := json.NewEncoder(c.writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.writer, err.Error(), 500)
	}
}

func (c *Context) HTML(code int, html string) {
	c.setHeader("Content-Type", "text/html")
	c.setCode(200)

	c.writer.Write([]byte(html))
}

func (c *Context) setCode(code int) {
	c.StatusCode = code
	c.writer.WriteHeader(code)
}

func (c *Context) setHeader(key, value string) {
	c.writer.Header().Set(key, value)
}
