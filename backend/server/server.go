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
	sr := &structs.SearchRequest{
		FromTimestamp: r.FormValue("fromTimestamp"),
		ToTimestamp:   r.FormValue("toTimestamp"),
		Name:          r.FormValue("name"),
		Query:         r.FormValue("query"),
	}
	if err := sr.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		makeError(w, Error{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("unable to validate request: %v", err),
		})
		return
	}
	result, err := gl.Search(sr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		makeError(w, Error{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("unable to search data: %v", err),
		})
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		makeError(w, Error{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("unable to marshal data: %v", err),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		makeError(w, Error{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("unable to write response: %v", err),
		})
		return
	}
}

// stats returns current statistics
func getServerStats(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(gl.Stat())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Fprintf(w, "%v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Fprint(w, "unable to write data")
		return
	}
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

	go func() {
		err := gl.SendEvent(toLogRequest(c))
		if err != nil {
			fmt.Printf("unable to send event: %v\n", err)
		}
	}()
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
	r.Get("/api/selitra/server", getServerStats)
	r.Post("/api/selitra/stats", postStats)
	fmt.Printf("Starting of the server at %s...\n", c.Address)
	fmt.Println(http.ListenAndServe(c.Address, r))
}
