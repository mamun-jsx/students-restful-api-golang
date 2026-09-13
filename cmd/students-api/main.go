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

	"github.com/mamun-jsx/students-restful-api-golang/internal/config"
	"github.com/mamun-jsx/students-restful-api-golang/internal/student"
)

func main() {
	// TODO: load config
	cfg := config.MustLoad()
	// TODO database setup

	// TODO setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /students", student.New())

	fmt.Println("welcome to student api")

	// TODO setup server
	server := http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: router,
	}
	fmt.Printf("Starting server on %s\n", cfg.HttpServer.Address)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("failed to start server")
		}
	}()
	<-done
	slog.Info("shutting down the server")
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown server ", slog.String("error", err.Error()))
	}
	slog.Info("server gracefully stopped")

}

// 7h: 38m
