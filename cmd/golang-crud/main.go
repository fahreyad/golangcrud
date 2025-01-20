package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fahreyad/golangcrud/internal/config"
)

func main() {
	fmt.Println("Hello, World!")
	//load config
	cfg := config.MustLoad()
	//database setup
	//set up router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to golang crud abcddd"))
	})
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
	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server: ", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown")
}
