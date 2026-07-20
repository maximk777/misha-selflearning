package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateAndListTask(t *testing.T) {
	h := NewHandler()
	create := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"learn Go"}`))
	create.Header.Set("Content-Type", "application/json")
	created := httptest.NewRecorder()
	h.ServeHTTP(created, create)
	if created.Code != http.StatusCreated || created.Header().Get("X-Request-ID") == "" {
		t.Fatalf("status=%d headers=%v", created.Code, created.Header())
	}
	listed := httptest.NewRecorder()
	h.ServeHTTP(listed, httptest.NewRequest(http.MethodGet, "/tasks", nil))
	if listed.Code != http.StatusOK || !strings.Contains(listed.Body.String(), "learn Go") {
		t.Fatalf("%d %s", listed.Code, listed.Body.String())
	}
}
func TestRejectsInvalidJSON(t *testing.T) {
	w := httptest.NewRecorder()
	NewHandler().ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{}`)))
	if w.Code != http.StatusBadRequest || !strings.Contains(w.Body.String(), `"error"`) {
		t.Fatalf("%d %s", w.Code, w.Body.String())
	}
}
