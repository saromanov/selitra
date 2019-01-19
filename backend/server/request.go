package server

// Request defines log payload
type Request struct {
	Level   uint   `json:"level"`
	Message string `json:"message"`
}
