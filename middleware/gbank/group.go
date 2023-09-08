package gbank

import (
	"html/template"
)

type RouteGroup struct {
	prefix string
	parent *RouteGroup
	engine *Engine

	// 中间件
	middleWares []HandlerFunc

	// 模板相关
	templates *template.Template
	funcMap   template.FuncMap
}

func (r *RouteGroup) NewGroup(prefix string) *RouteGroup {
	group := &RouteGroup{
		prefix: r.prefix + prefix,
		parent: r,
		engine: r.engine,
	}
	r.engine.routeGroups = append(r.engine.routeGroups, group)
	return group
}

// Use 添加中间件
func (r *RouteGroup) Use(handlerFunc ...HandlerFunc) {
	r.middleWares = append(r.middleWares, handlerFunc...)
}

func (r *RouteGroup) GET(path string, handlerFunc HandlerFunc) {
	r.addRoute("GET", path, handlerFunc)
}

func (r *RouteGroup) POST(path string, handlerFunc HandlerFunc) {
	r.addRoute("POST", path, handlerFunc)
}

func (r *RouteGroup) addRoute(method, path string, handlerFunc HandlerFunc) {
	pattern := r.prefix + path
	r.engine.router.AddRoute(method, pattern, handlerFunc)
}
