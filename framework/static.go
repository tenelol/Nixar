// framework/static.go
package framework

import "net/http"

// 静的ファイルを返すハンドラを返す
func Static(dir string) http.Handler {
	return http.FileServer(http.Dir(dir))
}

