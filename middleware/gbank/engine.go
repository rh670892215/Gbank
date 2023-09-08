package gbank

import (
	"net/http"
	"strings"
)

type Engine struct {
	RouteGroup
	routeGroups []*RouteGroup
	router      *Router
}

// Default 初始化engine，默认使用logger、recover中间件
func Default() *Engine {
	res := NewEngine()
	res.Use(Logger(), Recover())
	return res
}

func NewEngine() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouteGroup = RouteGroup{engine: engine}
	engine.routeGroups = []*RouteGroup{&engine.RouteGroup}
	return engine
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handles []HandlerFunc
	for _, group := range e.routeGroups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			handles = append(handles, group.middleWares...)
		}
	}
	ctx := NewContext(w, req)
	ctx.handles = handles
	e.router.Handle(ctx)
}

func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}
