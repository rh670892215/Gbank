package gbank

import (
	"html/template"
	"net/http"
	"path"
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

// Static 请求静态文件，relativePath是请求的相对路径 root是服务端文件根路径
func (r *RouteGroup) Static(relativePath string, root string) {
	handler := r.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")

	// 注册handler
	r.GET(urlPattern, handler)
}

// 创建文件请求handler
func (r *RouteGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(r.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))

	return func(c *Context) {
		fileName := c.params["filepath"]
		if _, err := fs.Open(fileName); err != nil {
			c.setCode(500)
			return
		}
		fileServer.ServeHTTP(c.writer, c.Req)
	}
}

// LoadHTMLTemplate 加载html模板文件
func (r *RouteGroup) LoadHTMLTemplate(pattern string) {
	r.templates = template.Must(template.New("").Funcs(r.funcMap).ParseGlob(pattern))
}

// SetFuncMap 设置html模板内嵌函数
func (r *RouteGroup) SetFuncMap(funcMap template.FuncMap) {
	r.funcMap = funcMap
}
