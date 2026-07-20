package main

//go:generate go run github.com/ogen-go/ogen/cmd/ogen@v1.16.0 --target internal/ogen --package ogen api/openapi.yaml

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/maximkirienkov/misha-backend-lab/internal/config"
	"github.com/maximkirienkov/misha-backend-lab/internal/platform/httpserver"
	"github.com/maximkirienkov/misha-backend-lab/internal/task"
)

func main() {
	config := config.Load()
	service := task.NewService(task.NewMemoryRepository(), &task.MemoryPublisher{}, task.NewMemoryCache(time.Minute))
	server := httpserver.New(config.Address, httpserver.NewHandler(service))
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	shutdown, cancel := context.WithTimeout(context.Background(), config.ShutdownTimeout)
	defer cancel()
	if err := server.Shutdown(shutdown); err != nil {
		log.Printf("shutdown: %v", err)
	}
}
