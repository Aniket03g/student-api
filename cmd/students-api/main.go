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

	"github.com/Aniket03g/students-api/internal/config"
	"github.com/Aniket03g/students-api/internal/config/http/handleres/student"
)

func main() {
	//load config
	cfg := config.Mustload()
	//database setup
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())
	//setup server

	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.HTTPServer.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Server Shutdown Failed", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown succesfully") //we did graceful shutdown here

}
