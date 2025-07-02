package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/sentientpixels/ava-bot-server/avacore"
)

func main() {
	fmt.Println("Loading avabot...")
	avacore.CheckEnv()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	r.Get("/", rootResponse)

	r.Post("/info", botInfo)

	fmt.Println("Bot loaded!")
	fmt.Println("Listening at port 3729...")
	http.ListenAndServe(":3729", r)
}

func botInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(avacore.GetBotInfo())
}

func rootResponse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}
