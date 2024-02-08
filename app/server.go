package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type app struct {
	server *http.Server
}

func (a app) Start() {
	err := a.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		slog.Info("Server is shutting down...")
	} else {
		slog.Error("Server could not be initiated...", err)
	}
}

func (a app) WaitToStop() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	<-ctx.Done()

	if err := a.server.Shutdown(context.Background()); err != nil {
		slog.Error("fail to stop the server properly...")
	}

	slog.Info("Server has been closed")
	stop()
}

func NewApp() *app {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Timeout(60 * time.Second))

	registerAppRoutes(router)

	return &app{
		server: &http.Server{
			Addr:         ":7001",
			Handler:      router,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}
