package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aerlaut/go-chess-backend/internal/engine"
	"github.com/aerlaut/go-chess-backend/internal/game"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

const PORT = "5000"

func main() {

	engine.InitEngine()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/api", func(r chi.Router) {
		// Upgrade to websocket connection
		r.Get("/game/{gameId}", game.ConnectToGame)
		r.Get("/game", game.GenerateGameLink)
	})

	log.Println("[SERVER] === Go-Chess backend server ===")
	log.Printf("[SERVER] - Listening on port %s", PORT)

	http.ListenAndServe(fmt.Sprintf("localhost:%s", PORT), r)
}
