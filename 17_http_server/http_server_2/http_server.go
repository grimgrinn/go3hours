package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("Request received, sleeping 5 seconds...")
	// 	time.Sleep(5 * time.Second)
	// 	w.Write([]byte("No graceful shutwodwn...\n"))
	// })

	// srv := &http.Server{Addr: ":8080", Handler: handler}

	// log.Println("Server started on :8000")
	// if err := srv.ListenAndServe(); err != nil {
	// 	log.Fatal(err)
	// }

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received, sleeping 5 seconds...")
		time.Sleep(5 * time.Second)
		w.Write([]byte("Hello, graceful shutwodwn...\n"))
	})

	srv := &http.Server{Addr: ":8080", Handler: handler}

	go func() {
		log.Println("Server started on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	log.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("Shutdown error: %v", err)
	}

	log.Println("Server gracefully stopped")
}
