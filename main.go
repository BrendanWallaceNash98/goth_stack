package main

import (
	"embed"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"goth_stack/handler"
	"log"
	"log/slog"
	"net/http"
	"os"
)

var FS embed.FS

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))
	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("application running", "port", port)
	log.Fatal(http.ListenAndServe(port, router))

}

func initEverything() error {
	return godotenv.Load()
}
