package server

// Request defines log payload
type Request struct {
	Level   string                 `json:"level"`
	Message string                 `json:"message"`
	Entries map[string]interface{} `json:"entries"`
	Name    string                 `json:"name"`
	Labels  []string               `json:"labels"`
}
