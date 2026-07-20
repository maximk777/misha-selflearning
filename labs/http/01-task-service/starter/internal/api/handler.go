package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
type errorEnvelope struct {
	Error string `json:"error"`
}
type service struct {
	mu     sync.Mutex
	nextID int
	tasks  []Task
}

func NewHandler() http.Handler {
	s := &service{nextID: 1}
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", s.tasksHandler)
	return requestID(mux)
}

func (s *service) tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.mu.Lock()
		tasks := append([]Task(nil), s.tasks...)
		s.mu.Unlock()
		writeJSON(w, http.StatusOK, tasks)
	case http.MethodPost:
		defer r.Body.Close()
		decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
		decoder.DisallowUnknownFields()
		var input struct {
			Title string `json:"title"`
		}
		if err := decoder.Decode(&input); err != nil || strings.TrimSpace(input.Title) == "" {
			writeJSON(w, http.StatusBadRequest, errorEnvelope{Error: "title is required"})
			return
		}
		select {
		case <-r.Context().Done():
			writeJSON(w, http.StatusRequestTimeout, errorEnvelope{Error: r.Context().Err().Error()})
			return
		default:
		}
		s.mu.Lock()
		task := Task{ID: s.nextID, Title: strings.TrimSpace(input.Title)}
		s.nextID++
		s.tasks = append(s.tasks, task)
		s.mu.Unlock()
		writeJSON(w, http.StatusCreated, task)
	default:
		w.Header().Set("Allow", "GET, POST")
		writeJSON(w, http.StatusMethodNotAllowed, errorEnvelope{Error: "method not allowed"})
	}
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
