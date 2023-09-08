package gbank

import "strings"

type Router struct {
	// 前缀树根节点
	root map[string]*Node
	// method + path 与 HandlerFunc 的映射关系
	handlerFuncMap map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		root:           make(map[string]*Node),
		handlerFuncMap: make(map[string]HandlerFunc, 0),
	}
}

// AddRoute 添加路由规则
func (r *Router) AddRoute(method, path string, handlerFunc HandlerFunc) {
	key := method + "_" + path
	r.handlerFuncMap[key] = handlerFunc

	parts := r.parsePath(path)
	_, ok := r.root[method]
	if !ok {
		// 若没有对应method的根节点，则先创建根节点
		r.root[method] = &Node{}
	}
	// 插入新节点
	r.root[method].Insert(parts, path, 0)
}

// 解析path，只允许一个*
func (r *Router) parsePath(path string) []string {
	parts := strings.Split(path, "/")
	var res []string
	for _, part := range parts {
		if part == "" {
			continue
		}
		res = append(res, part)

		if strings.HasPrefix(part, "*") {
			break
		}
	}
	return res
}

// Handle 处理请求
func (r *Router) Handle(c *Context) {
	n, params := r.getRoute(c.method, c.path)
	if n == nil {
		c.String(404, "not found path %s", c.path)
		return
	}
	c.params = params
	key := c.method + "_" + n.pattern
	handle, ok := r.handlerFuncMap[key]
	if !ok {
		c.String(404, "not found path %s", c.path)
		return
	}

	handle(c)
}

// 根据method、path查找对应的节点，返回查找的节点和映射关系
// 如 /hello/bank - /hello/:name 		=> name - bank
// 如 /get/css/log.txt - /get/*filepath => filepath - css/log.txt
func (r *Router) getRoute(method, path string) (*Node, map[string]string) {
	rootNode, ok := r.root[method]
	if !ok {
		return nil, nil
	}

	parts := r.parsePath(path)
	n := rootNode.Search(parts, 0)
	if n == nil {
		return nil, nil
	}
	searchRes := r.parsePath(n.pattern)
	res := make(map[string]string, 0)
	for index, part := range searchRes {
		if part[0] == ':' {
			res[part[1:]] = parts[index]
		}
		// 若匹配到 * ，则取后面所有的parts
		if part[0] == '*' && len(part) > 1 {
			res[part[1:]] = strings.Join(parts[index:], "/")
			break
		}
	}
	return n, res
}
