package v1

// LogRequest provides getting of data from logrus
type LogRequest struct {
	Level     string
	Message   string
	Entry     string `json:"entry"`
	Service   string `json:"service"`
	Name      string
	Labels    []string
	Timestamp int64
}

// SearchRequest defines struct for searching
type SearchRequest struct {
	FromTimestamp int64
	ToTimestamp   int64
}
