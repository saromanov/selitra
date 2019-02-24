package v1

import (
	"fmt"
	"time"
)

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
	FromTimestamp string
	ToTimestamp   string
	Name          string
	Query         string
	FromDate string
	ToDate string
}

// Validate provides validation of the input for searchrequest
func (s *SearchRequest) Validate() error {
	if s.FromDate != "" {
		_, err := time.Parse(time.RFC3339, s.FromDate)
		if err != nil {
			return fmt.Errorf("fromDate should be on RFC 3339 format: %v", err)
		}
	}

	if s.ToDate != "" {
		_, err := time.Parse(time.RFC3339, s.ToDate)
		if err != nil {
			return fmt.Errorf("toDate should be on RFC 3339 format: %v", err)
		}
	}

	return nil
}
