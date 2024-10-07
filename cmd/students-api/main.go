package main

import (
	"context"
	"fmt"
	"github.com/Arunshaik2001/students-api/internal/config"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.LoadConfig()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Welcome to Students API"))
		if err != nil {
			log.Fatalf("Error writing response %s", err)
		}
	})

	fmt.Printf("Starting students-api %s\n", cfg)

	slog.Info("server started", slog.String("address", cfg.HttpServer.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	server := http.Server{
		Addr:    cfg.HttpServer.Addr,
		Handler: router,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Error writing response %s", err)
		}
	}()

	<-done

	fmt.Println("Shutting down students-api")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down students-api %s", err)
	}
}
