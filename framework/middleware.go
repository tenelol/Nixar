package framework

import (
	"log"
	"time"
)

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
