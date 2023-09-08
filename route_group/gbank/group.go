package gbank

type RouteGroup struct {
	prefix string
	parent *RouteGroup
	engine *Engine
}

func (r *RouteGroup) NewGroup(prefix string) *RouteGroup {
	group := &RouteGroup{
		prefix: r.prefix + prefix,
		parent: r,
		engine: r.engine,
	}
	return group
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
