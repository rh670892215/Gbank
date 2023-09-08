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

	// 存储前缀树解析路径结果
	params map[string]string

	index   int
	handles []HandlerFunc

	engine *Engine
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		writer: w,
		Req:    r,
		method: r.Method,
		path:   r.URL.Path,
		index:  -1,
	}
}

// Next 将执行权交给下一个中间件
func (c *Context) Next() {
	c.index++
	for ; c.index < len(c.handles); c.index++ {
		c.handles[c.index](c)
	}
}

// Fail 服务终止
func (c *Context) Fail(message string) {
	c.index = len(c.handles)
	http.Error(c.writer, message, 500)
}

// QueryGet 查询get请求行中的参数
func (c *Context) QueryGet(key string) string {
	return c.Req.URL.Query().Get(key)
}

// QueryPost 查询post请求体中数据
func (c *Context) QueryPost(key string) string {
	return c.Req.PostForm.Get(key)
}

// String string类型返回数据
func (c *Context) String(code int, format string, params ...string) {
	c.setCode(code)
	c.writer.Write([]byte(fmt.Sprintf(format, params)))
}

// JSON json类型返回数据
func (c *Context) JSON(code int, obj interface{}) {
	c.setHeader("Content-Type", "application/json")
	c.setCode(code)

	encoder := json.NewEncoder(c.writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.writer, err.Error(), 500)
	}
}

// HTML html类型返回数据
func (c *Context) HTML(code int, name string, data interface{}) {
	c.setHeader("Content-Type", "text/html")
	c.setCode(code)

	if err := c.engine.templates.ExecuteTemplate(c.writer, name, data); err != nil {
		http.Error(c.writer, err.Error(), 500)
	}
}

// 设置连接状态码
func (c *Context) setCode(code int) {
	c.StatusCode = code
	c.writer.WriteHeader(code)
}

// 设置请求头
func (c *Context) setHeader(key, value string) {
	c.writer.Header().Set(key, value)
}

// GetParam 获取前缀树的路径解析结果
func (c *Context) GetParam(key string) string {
	res, _ := c.params[key]
	return res
}
