package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received, sleeping 5 seconds...")
		time.Sleep(5 * time.Second)
		w.Write([]byte("No graceful shutwodwn...\n"))
	})

	srv := &http.Server{Addr: ":8080", Handler: handler}

	log.Println("Server started on :8000")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
