package app

// LogRequest provides getting of data from logrus
type LogRequest struct {
	Level   string
	Message string
	Entries map[string]interface{}
	Name    string
	Labels  []string
}
