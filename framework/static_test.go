package framework

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestStaticNoCacheHeaders(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "app.css")
	if err := os.WriteFile(filePath, []byte("body{}"), 0644); err != nil {
		t.Fatalf("write file: %v", err)
	}

	handler := Static(dir)
	req := httptest.NewRequest(http.MethodGet, "/app.css", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if got := rr.Header().Get("Cache-Control"); got == "" {
		t.Fatalf("expected Cache-Control header to be set")
	}
	if got := rr.Header().Get("Pragma"); got == "" {
		t.Fatalf("expected Pragma header to be set")
	}
	if got := rr.Header().Get("Expires"); got == "" {
		t.Fatalf("expected Expires header to be set")
	}
}
