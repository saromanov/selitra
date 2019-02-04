package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	err := decoder.Decode(&c)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "unable to decode data: %v", err)
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
		Level:     r.Level,
		Message:   r.Message,
		Entry:     r.Entry,
		Name:      r.Name,
		Labels:    r.Labels,
		Timestamp: time.Now().UnixNano(),
	}
}

// Create provides initialization of server
func Create(a *app.App, c *Config) {
	gl = a
	r := chi.NewRouter()
	r.Get("/api/selitra/stats", stats)
	r.Post("/api/selitra/stats", postStats)
	fmt.Println("Starting of the server...")
	fmt.Println(http.ListenAndServe(c.Address, r))
}
