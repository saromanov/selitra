package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error defines the struct for response on the error case
type Error struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Additional string `json:"additional"`
}

func makeError(w http.ResponseWriter, e Error) {
	res, _ := json.Marshal(e)
	fmt.Fprint(w, res)
}
