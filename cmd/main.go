package main

import (
	"fmt"
	"net/http"

	"github.com/alessandra1408/crud-golang/configs"
	"github.com/alessandra1408/crud-golang/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {

	err := configs.Load()
	if err != nil {
		panic(fmt.Sprintf("Error to load variables: %v", err))
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.Get)
	r.Get("/{id}", handlers.List)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
