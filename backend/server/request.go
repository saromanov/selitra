package server

// Request defines log payload
type Request struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}
