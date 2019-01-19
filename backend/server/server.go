package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/saromanov/selitra/backend/internal/app"
)

var a *app.App

// stats returns current statistics
func stats(w http.ResponseWriter, r *http.Request) {

}

func postStats(w http.ResponseWriter, r *http.Request) {
	var c *Request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(c)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	go a.SendEvent(&app.LogRequest{
		Level: c.Level,
	})
	w.WriteHeader(http.StatusCreated)

}

// Create provides initialization of server
func Create(a *app.App, c *Config) {
	a = a
	r := chi.NewRouter()
	r.Get("/stats", stats)
	r.Post("/stats", postStats)
	http.ListenAndServe(c.Address, r)
}
