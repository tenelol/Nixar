package framework

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

//
// ===== 基本の型 =====
//

type Context struct {
	W      http.ResponseWriter
	Req    *http.Request
	Params map[string]string
	App    *App
}

type HandlerFunc func(*Context)
type Middleware func(HandlerFunc) HandlerFunc

//
// ===== Router 部分 =====
//

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

// App から呼ばれる内部用ルータ
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

	if len(pSegs) != len(pathSegs) {
		return nil, false
	}

	params := make(map[string]string)

	for i := 0; i < len(pSegs); i++ {
		pp := pSegs[i]
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

//
// ===== App 本体 =====
//

type App struct {
	router      *Router
	middlewares []Middleware

	Logger *log.Logger
}

func NewApp() *App {
	return &App{
		router:      NewRouter(),
		middlewares: []Middleware{},
		Logger:      log.Default(),
	}
}

func (a *App) Use(m Middleware) {
	a.middlewares = append(a.middlewares, m)
}

func (a *App) Handle(method, pattern string, h HandlerFunc) {
	a.router.Handle(method, pattern, h)
}

func (a *App) Get(pattern string, h HandlerFunc) {
	a.Handle(http.MethodGet, pattern, h)
}

func (a *App) Post(pattern string, h HandlerFunc) {
	a.Handle(http.MethodPost, pattern, h)
}

func (a *App) Put(pattern string, h HandlerFunc) {
	a.Handle(http.MethodPut, pattern, h)
}

func (a *App) Delete(pattern string, h HandlerFunc) {
	a.Handle(http.MethodDelete, pattern, h)
}

// http.Handler 実装
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Router を呼ぶための HandlerFunc
	base := func(ctx *Context) {
		a.router.serve(ctx)
	}

	// ミドルウェアを内側からラップしてパイプラインを組む
	h := base
	for i := len(a.middlewares) - 1; i >= 0; i-- {
		h = a.middlewares[i](h)
	}

	// Context 作成
	ctx := &Context{
		W:      w,
		Req:    r,
		Params: map[string]string{},
		App:    a,
	}

	// 実行
	h(ctx)
}

//
// ===== Context のヘルパー =====
//

func (c *Context) Param(key string) string {
	if c.Params == nil {
		return ""
	}
	return c.Params[key]
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) QueryDefault(key, def string) string {
	v := c.Query(key)
	if v == "" {
		return def
	}
	return v
}

func (c *Context) JSON(status int, v any) {
	JSON(c, status, v)
}

func (c *Context) HTMLFile(path string) {
	c.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.W.Header().Set("Pragma", "no-cache")
	c.W.Header().Set("Expires", "0")

	http.ServeFile(c.W, c.Req, path)
}

//
// ===== ミドルウェア =====
//

func Logging() Middleware {
	return func(next HandlerFunc) HandlerFunc {
		return func(c *Context) {
			start := time.Now()

			next(c)

			duration := time.Since(start)

			if c.App != nil && c.App.Logger != nil {
				c.App.Logger.Printf("%s %s %s", c.Req.Method, c.Req.URL.Path, duration)
			} else {
				log.Printf("%s %s %s", c.Req.Method, c.Req.URL.Path, duration)
			}
		}
	}
}

//
// ===== レスポンスユーティリティ =====
//

func JSON(ctx *Context, status int, data any) {
	ctx.W.Header().Set("Content-Type", "application/json; charset=utf-8")
	ctx.W.WriteHeader(status)
	_ = json.NewEncoder(ctx.W).Encode(data)
}

//
// ===== 静的ファイル／ラッパ =====
//

func Static(dir string) http.Handler {
	fs := http.FileServer(http.Dir(dir))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		fs.ServeHTTP(w, r)
	})
}

func WrapHTTPHandler(h http.Handler) HandlerFunc {
	return func(c *Context) {
		h.ServeHTTP(c.W, c.Req)
	}
}

func WrapHTTPHandlerFunc(h http.HandlerFunc) HandlerFunc {
	return func(c *Context) {
		h(c.W, c.Req)
	}
}

