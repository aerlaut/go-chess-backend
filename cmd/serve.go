package main

import (
	"fmt"
	"net/http"

	"github.com/aerlaut/go-chess-backend/pkg/matcher"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const PORT = "5000"

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", matcher.GenerateMatchLink)

	fmt.Println("=== Go-Chess backend server ===")
	fmt.Printf("Listening on port %s", PORT)

	http.ListenAndServe(fmt.Sprintf("localhost:%s", PORT), r)
}
