package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// LogRequest defines model for log storage
type LogRequest struct {
	gorm.Model
	Message   string
	Name      string
	Labels    pq.StringArray `gorm:"type:varchar(64)[]"`
	Timestamp uint64
	Entry     string
	Service   string
}

// SearchRequest defines request for getting logs
type SearchRequest struct {
	FromTimestamp int64
	ToTimestamp   int64
	Name          string
	Query         string
}
