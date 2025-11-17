// framework/router.go
package framework

import (
	"net/http"
	"strings"
)

type Context struct {
	W      http.ResponseWriter
	Req    *http.Request
	Params map[string]string
}

type HandlerFunc func(*Context)

type route struct {
	method  string
	pattern string
	handler HandlerFunc
}

type Router struct {
	routes []route
}

func NewRouter() *Router {
	return &Router{
		routes: make([]route, 0),
	}
}

// e.g. Handle("GET", "/users/:id", handler)
func (r *Router) Handle(method, pattern string, handler HandlerFunc) {
	r.routes = append(r.routes, route{
		method:  method,
		pattern: pattern,
		handler: handler,
	})
}

// http.Server から呼ばれる入口
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path

	for _, rt := range r.routes {
		if rt.method != method {
			continue
		}
		if params, ok := matchPattern(rt.pattern, path); ok {
			ctx := &Context{
				W:      w,
				Req:    req,
				Params: params,
			}
			rt.handler(ctx)
			return
		}
	}

	http.NotFound(w, req)
}

// pattern: "/users/:id"
// path:    "/users/123"
func matchPattern(pattern, path string) (map[string]string, bool) {
	pSegs := splitPath(pattern)
	pathSegs := splitPath(path)

	if len(pSegs) != len(pathSegs) {
		return nil, false
	}

	params := make(map[string]string)

	for i := 0; i < len(pSegs); i++ {
		pp := pSegs[i]
		ps := pathSegs[i]

		if strings.HasPrefix(pp, ":") {
			// パラメータ
			key := pp[1:]
			params[key] = ps
			continue
		}

		// 通常の文字列は一致必須
		if pp != ps {
			return nil, false
		}
	}

	return params, true
}

func splitPath(p string) []string {
	p = strings.TrimSpace(p)
	if p == "" {
		return []string{}
	}
	// 両端の / を削る
	p = strings.TrimPrefix(p, "/")
	p = strings.TrimSuffix(p, "/")
	if p == "" {
		return []string{}
	}
	return strings.Split(p, "/")
}

