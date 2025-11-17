// framework/middleware.go
package framework

import (
	"log"
	"net/http"
	"time"
)

// http.Handler をラップする形のシンプルなミドルウェア
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 本来の処理
		next.ServeHTTP(w, r)

		// 終わった後にログ
		duration := time.Since(start)
		log.Printf("%s %s %s", r.Method, r.URL.Path, duration)
	})
}

