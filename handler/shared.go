package handler

import (
	"log/slog"
	"net/http"
)

func MakeHandler(h func(w http.ResponseWriter, response *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("internal server err", "err", err, "path", r.URL.Path)
		}
	}
}
