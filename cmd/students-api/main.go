package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mamun-jsx/students-restful-api-golang/internal/config"
)

func main() {
	// TODO: load config
	cfg := config.MustLoad()
	// TODO database setup

	// TODO setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students api"))
	})

	fmt.Println("welcome to student api")

	// TODO setup server
	server := http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: router,
	}
	fmt.Printf("Starting server on %s\n", cfg.HttpServer.Address)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start server")
	}

}
