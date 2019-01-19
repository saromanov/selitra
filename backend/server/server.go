package server

import (
	"net/http"
	"encoding/json"

	"github.com/go-chi/chi"
	"github.com/saromanov/selitra/backend/internal/app"
)

var a *app.App

// stats returns current statistics
func stats(w http.ResponseWriter, r *http.Request) {

}

func postStats(w http.ResponseWriter, r *http.Request) {
	var c *app.Request
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(messageObject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonapi.WriteBasicError(w, "400", err.Error(), "Cannot messageObject")
		return
	}
	
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
