package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	// router.Get("/",) {
	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("application running", "port", port)
	log.Fatal(http.ListenAndServe(port, router))

}

func initEverything() error {
	return godotenv.Load()
}
