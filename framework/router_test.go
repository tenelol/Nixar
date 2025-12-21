package framework

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterParamMatch(t *testing.T) {
	app := NewApp()
	app.Get("/users/:id", func(ctx *Context) {
		ctx.W.WriteHeader(http.StatusOK)
		_, _ = ctx.W.Write([]byte("id=" + ctx.Param("id")))
	})

	req := httptest.NewRequest(http.MethodGet, "/users/42", nil)
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if body := rr.Body.String(); body != "id=42" {
		t.Fatalf("unexpected body: %q", body)
	}
}

func TestRouterWildcardMatch(t *testing.T) {
	app := NewApp()
	app.Get("/static/*file", func(ctx *Context) {
		ctx.W.WriteHeader(http.StatusOK)
		_, _ = ctx.W.Write([]byte(ctx.Param("file")))
	})

	req := httptest.NewRequest(http.MethodGet, "/static/css/app.css", nil)
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if body := rr.Body.String(); body != "css/app.css" {
		t.Fatalf("unexpected body: %q", body)
	}
}
