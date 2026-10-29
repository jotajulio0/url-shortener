package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		os.Exit(1)
	}
	slog.Info("All Systems offline")
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/url", handler.ShortenURL)
	r.Get("/url/{code}", handler.RedirectURL)

	server := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           r,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    16 * 1024,
	}

	err = http.ListenAndServe(server.Addr, server.Handler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server started on port " + server.Addr)
	return nil
}
