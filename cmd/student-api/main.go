package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fahreyad/golangcrud/internal/config"
	"github.com/fahreyad/golangcrud/internal/http/handlers/student"
	"github.com/fahreyad/golangcrud/internal/storage/sqlite"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	slog.Info("database connected", slog.String("storage_path", cfg.StoragePath), slog.String("env", cfg.ENV))

	//set up router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))
	//set up server
	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}
	slog.Info("server started", slog.String("address", cfg.HTTPServer.Address))
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server: ", err)
		}
	}()

	<-done

	slog.Info("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server: ", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown")
}
