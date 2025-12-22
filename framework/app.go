package framework

import (
	"log"
	"net/http"
)

type HandlerFunc func(*Context)
type Middleware func(HandlerFunc) HandlerFunc

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

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := func(ctx *Context) {
		a.router.serve(ctx)
	}

	h := base
	for i := len(a.middlewares) - 1; i >= 0; i-- {
		h = a.middlewares[i](h)
	}

	ctx := &Context{
		W:      w,
		Req:    r,
		Params: map[string]string{},
		App:    a,
	}

	h(ctx)
}
