package server

// Request defines log payload
type Request struct {
	Level   string   `json:"level"`
	Message string   `json:"message"`
	Name    string   `json:"name"`
	Labels  []string `json:"labels"`
	Entry   string   `json:"entry"`
	Service string   `json:"service"`
}
