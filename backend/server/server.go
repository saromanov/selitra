package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/saromanov/selitra/backend/internal/app"
	structs "github.com/saromanov/selitra/backend/internal/structs/v1"
)

var gl *app.App

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
	if c == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	go gl.SendEvent(toLogRequest(c))
	w.WriteHeader(http.StatusCreated)

}

// toLogRequest convert request to inner representation
func toLogRequest(r *Request) *structs.LogRequest {
	return &structs.LogRequest{
		Level:   r.Level,
		Message: r.Message,
		Entries: r.Entries,
		Name:    r.Name,
		Labels:  r.Labels,
	}
}

// Create provides initialization of server
func Create(a *app.App, c *Config) {
	gl = a
	r := chi.NewRouter()
	r.Get("/api/selitra/stats", stats)
	r.Post("/api/selitra/stats", postStats)
	http.ListenAndServe(c.Address, r)
}
