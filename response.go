package framework

import (
	"encoding/json"
)

func JSON(ctx *Context, status int, data any) {
	ctx.W.Header().Set("Content-Type", "application/json; charset=utf-8")
	ctx.W.WriteHeader(status)
	_ = json.NewEncoder(ctx.W).Encode(data)
}

