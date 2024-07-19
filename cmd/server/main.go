package main

import (
	"net/http"

	"github.com/brenddonanjos/multithreading_api/internal/webserver/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	route := chi.NewRouter()

	cepHandler := handlers.NewCepHandler()
	route.Get("/{cep}", cepHandler.GetCepInfo)

	http.ListenAndServe(":8000", route)
}
