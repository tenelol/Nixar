package framework

import (
	"net/http"
	"strings"
)

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

func (r *Router) Handle(method, pattern string, handler HandlerFunc) {
	r.routes = append(r.routes, route{
		method:  method,
		pattern: pattern,
		handler: handler,
	})
}

func (r *Router) serve(ctx *Context) {
	method := ctx.Req.Method
	path := ctx.Req.URL.Path

	for _, rt := range r.routes {
		if rt.method != method {
			continue
		}
		if params, ok := matchPattern(rt.pattern, path); ok {
			ctx.Params = params
			rt.handler(ctx)
			return
		}
	}

	http.NotFound(ctx.W, ctx.Req)
}

func matchPattern(pattern, path string) (map[string]string, bool) {
	pSegs := splitPath(pattern)
	pathSegs := splitPath(path)

	params := make(map[string]string)

	for i := 0; i < len(pSegs); i++ {
		pp := pSegs[i]
		if strings.HasPrefix(pp, "*") {
			key := pp[1:]
			if len(pathSegs) < i {
				return nil, false
			}
			params[key] = strings.Join(pathSegs[i:], "/")
			return params, true
		}

		if len(pathSegs) <= i {
			return nil, false
		}
		ps := pathSegs[i]

		if strings.HasPrefix(pp, ":") {
			key := pp[1:]
			params[key] = ps
			continue
		}
		if pp != ps {
			return nil, false
		}
	}

	if len(pathSegs) != len(pSegs) {
		return nil, false
	}

	return params, true
}

func splitPath(p string) []string {
	p = strings.TrimSpace(p)
	if p == "" {
		return []string{}
	}
	p = strings.TrimPrefix(p, "/")
	p = strings.TrimSuffix(p, "/")
	if p == "" {
		return []string{}
	}
	return strings.Split(p, "/")
}
