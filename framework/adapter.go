package framework

import "net/http"
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
