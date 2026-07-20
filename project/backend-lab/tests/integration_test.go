package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maximkirienkov/misha-backend-lab/internal/platform/httpserver"
	"github.com/maximkirienkov/misha-backend-lab/internal/task"
)

func TestHealthDoesNotNeedDocker(t *testing.T) {
	handler := httpserver.NewHandler(task.NewService(task.NewMemoryRepository(), &task.MemoryPublisher{}))
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/readyz", nil))
	if response.Code != http.StatusOK {
		t.Fatalf("status=%d", response.Code)
	}
}
