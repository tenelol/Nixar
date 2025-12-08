package framework

import (
	"net/http"
)

type Context struct {
	W      http.ResponseWriter
	Req    *http.Request
	Params map[string]string
	App    *App
}

// ルートパラメータ
func (c *Context) Param(key string) string {
	if c.Params == nil {
		return ""
	}
	return c.Params[key]
}

// クエリ取得
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

// JSONレスポンス
func (c *Context) JSON(status int, v any) {
	JSON(c, status, v)
}

// HTMLファイル送信（no-cache）
func (c *Context) HTMLFile(path string) {
	c.W.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.W.Header().Set("Pragma", "no-cache")
	c.W.Header().Set("Expires", "0")

	http.ServeFile(c.W, c.Req, path)
}

