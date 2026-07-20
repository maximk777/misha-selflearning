package httpserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/maximkirienkov/misha-backend-lab/internal/task"
)

type Handler struct {
	service   *task.Service
	requestID atomic.Uint64
}

func NewHandler(service *task.Service) *Handler { return &Handler{service: service} }

func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	id := request.Header.Get("X-Request-ID")
	if id == "" {
		id = "req-" + strconvFormat(handler.requestID.Add(1))
	}
	writer.Header().Set("X-Request-ID", id)
	if request.URL.Path == "/healthz" || request.URL.Path == "/readyz" {
		writeJSON(writer, http.StatusOK, map[string]string{"status": "ok"})
		return
	}
	if request.Method == http.MethodPost && request.URL.Path == "/tasks" {
		handler.create(writer, request)
		return
	}
	if strings.HasPrefix(request.URL.Path, "/tasks/") {
		parts := strings.Split(strings.TrimPrefix(request.URL.Path, "/tasks/"), "/")
		if len(parts) == 1 && request.Method == http.MethodGet {
			handler.get(writer, request, parts[0])
			return
		}
		if len(parts) == 2 && parts[1] == "complete" && request.Method == http.MethodPost {
			handler.complete(writer, request, parts[0])
			return
		}
	}
	writeError(writer, http.StatusNotFound, "not_found")
}

func (handler *Handler) create(writer http.ResponseWriter, request *http.Request) {
	var input struct {
		Title          string `json:"title"`
		IdempotencyKey string `json:"idempotency_key"`
	}
	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		writeError(writer, http.StatusBadRequest, "invalid_json")
		return
	}
	value, created, err := handler.service.Create(request.Context(), input.Title, input.IdempotencyKey)
	if err != nil {
		writeTaskError(writer, err)
		return
	}
	status := http.StatusOK
	if created {
		status = http.StatusCreated
	}
	writeJSON(writer, status, value)
}

func (handler *Handler) get(writer http.ResponseWriter, request *http.Request, id string) {
	value, err := handler.service.Get(request.Context(), id)
	if err != nil {
		writeTaskError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, value)
}
func (handler *Handler) complete(writer http.ResponseWriter, request *http.Request, id string) {
	value, err := handler.service.Complete(request.Context(), id)
	if err != nil {
		writeTaskError(writer, err)
		return
	}
	writeJSON(writer, http.StatusOK, value)
}
func writeTaskError(writer http.ResponseWriter, err error) {
	if errors.Is(err, task.ErrNotFound) {
		writeError(writer, http.StatusNotFound, "not_found")
		return
	}
	if errors.Is(err, task.ErrConflict) {
		writeError(writer, http.StatusConflict, "conflict")
		return
	}
	if errors.Is(err, task.ErrInvalidTitle) {
		writeError(writer, http.StatusBadRequest, "validation_error")
		return
	}
	writeError(writer, http.StatusInternalServerError, "internal_error")
}
func writeError(writer http.ResponseWriter, status int, code string) {
	writeJSON(writer, status, map[string]string{"error": code})
}
func writeJSON(writer http.ResponseWriter, status int, value any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	_ = json.NewEncoder(writer).Encode(value)
}
func strconvFormat(value uint64) string { return strconv.FormatUint(value, 10) }
