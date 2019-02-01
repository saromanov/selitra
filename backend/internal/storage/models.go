package storage

import "github.com/jinzhu/gorm"

// LogRequest defines model for log storage
type LogRequest struct {
	gorm.Model
	Message string
	Entries map[string]interface{}
	Name    string
	Labels  []string
}
