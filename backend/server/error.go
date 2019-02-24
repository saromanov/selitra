package server

// Error defines the struct for response on the error case
type Error struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Additional string `json:"additional"`
}
