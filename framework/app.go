package framework

import (
	"log"
	"net/http"
)

// ハンドラとミドルウェアの型
type HandlerFunc func(*Context)
type Middleware func(HandlerFunc) HandlerFunc

// アプリ全体の状態を持つ構造体
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

// ミドルウェア追加
func (a *App) Use(m Middleware) {
	a.middlewares = append(a.middlewares, m)
}

// ルーティングAPI
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

	// ミドルウェアをラップしてパイプラインを作る
	h := base
	for i := len(a.middlewares) - 1; i >= 0; i-- {
		h = a.middlewares[i](h)
	}

	// Context を作成
	ctx := &Context{
		W:      w,
		Req:    r,
		Params: map[string]string{},
		App:    a,
	}

	// 実行
	h(ctx)
}

