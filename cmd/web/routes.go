package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (pp *application) routes() http.Handler {
	mux := chi.NewRouter()

	return mux
}
