package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	rec := httptest.NewRecorder()

	statusHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status code = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), `"message":"ok"`) {
		t.Fatalf("response body = %q", rec.Body.String())
	}
}

func TestLoggingMiddleware(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler := loggingMiddleware(http.HandlerFunc(statusHandler))
	handler.ServeHTTP(rec, req)

	if rec.Header().Get("X-Learn") != "golang" {
		t.Fatalf("X-Learn header = %q", rec.Header().Get("X-Learn"))
	}
}
