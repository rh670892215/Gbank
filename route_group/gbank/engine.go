package gbank

import "net/http"

type Engine struct {
	RouteGroup
	routeGroups []*RouteGroup
	router      *Router
}

func NewEngine() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouteGroup = RouteGroup{engine: engine}
	engine.routeGroups = []*RouteGroup{&engine.RouteGroup}
	return engine
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := NewContext(w, req)
	e.router.Handle(ctx)
}

func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}
