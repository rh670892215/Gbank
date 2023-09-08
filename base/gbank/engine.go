package gbank

import "net/http"

type Engine struct {
	router *Router
}

func NewEngine() *Engine {
	return &Engine{
		router: NewRouter(),
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := NewContext(w, req)
	e.router.Handle(ctx)
}

func (e *Engine) Get(path string, handle HandlerFunc) {
	e.router.AddRoute("GET", path, handle)
}

func (e *Engine) POST(path string, handle HandlerFunc) {
	e.router.AddRoute("POST", path, handle)
}

func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}
