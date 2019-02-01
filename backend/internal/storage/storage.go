// Package storage provides implementation of the storage for log metrics
package storage

// Storage defines interface for storage handling
type Storage interface {
	Insert(*LogRequest) error
	Close() error
}
