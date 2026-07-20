package client

import (
	"context"
	"github.com/maximkirienkov/misha-wzorvaniy/http-task-service/internal/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {
	srv := httptest.NewServer(api.NewHandler())
	defer srv.Close()
	got, err := New().List(context.Background(), srv.URL)
	if err != nil || len(got) != 0 {
		t.Fatalf("%v %v", got, err)
	}
}
func TestListReturnsStatusError(t *testing.T) {
	srv := httptest.NewServer(http.NotFoundHandler())
	defer srv.Close()
	if _, err := New().List(context.Background(), srv.URL); err == nil {
		t.Fatal("want status error")
	}
}
