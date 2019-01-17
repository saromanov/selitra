package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/saromanov/selitra/backend/internal/app"
)

// stats returns current statistics
func stats(w http.ResponseWriter, r *http.Request) {

}

func postStats(w http.ResponseWriter, r *http.Request) {

}

// Create provides initialization of server
func Create(a *app.App, c *Config) {
	r := chi.NewRouter()
	r.Get("/stats", stats)
	r.Post("/stats", postStats)
	http.ListenAndServe(c.Address, r)
}
