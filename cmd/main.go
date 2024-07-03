package main

import (
	"GoTTH-commerce/internal/handlers"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Timeout(60 * time.Second))

	r.Group(func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.RealIP)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)

		r.Get("/", handlers.NewHomeHandler().ServeHTTP)
	})

	err := http.ListenAndServe(":3333", r)
	if err != nil {
		fmt.Println(err)
	}
}
