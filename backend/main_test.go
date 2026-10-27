package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	home(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if rec.Body.String() != "GitLab CI/CD Backend Version 1.0" {
		t.Fatalf("unexpected body: %s", rec.Body.String())
	}
}

func TestHealthHandlerDatabaseDown(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler("127.0.0.1").ServeHTTP(rec, req)

	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected 503, got %d", rec.Code)
	}
}
