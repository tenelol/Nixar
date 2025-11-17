// framework/response.go
package framework

import (
	"encoding/json"
	"net/http"
)

func JSON(ctx *Context, status int, data any) {
	ctx.W.Header().Set("Content-Type", "application/json; charset=utf-8")
	ctx.W.WriteHeader(status)
	_ = json.NewEncoder(ctx.W).Encode(data)
}

// たまに生で使いたい時用に残してもいい
func JSONRaw(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

