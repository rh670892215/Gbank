package gbank

type Router struct {
	table map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		table: make(map[string]HandlerFunc, 0),
	}
}

func (r *Router) AddRoute(method, path string, handlerFunc HandlerFunc) {
	key := method + "_" + path
	r.table[key] = handlerFunc
}

func (r *Router) Handle(c *Context) {
	key := c.method + "_" + c.path
	handle, ok := r.table[key]
	if !ok {
		c.String(404, "not found path %s", c.path)
		return
	}

	handle(c)
}
