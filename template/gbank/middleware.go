package gbank

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

// Logger 日志插件
func Logger() HandlerFunc {
	return func(c *Context) {
		now := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(now))
	}
}

// Recover 错误恢复插件
func Recover() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, false)
				stackInfo := fmt.Sprintf("%s", buf[:n])
				// 打印调用栈信息
				log.Printf("%s", stackInfo)
				c.Fail("Internal Server Error")
			}
		}()

		c.Next()
	}
}
