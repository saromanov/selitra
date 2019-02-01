package postgresql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	st "github.com/saromanov/selitra/backend/internal/storage"
)

// storage implements db handling with Postgesql
type storage struct {
	db *gorm.DB
}

// Create provides init for postgesql storage
func Create() (st.Storage, error) {
	db, err := gorm.Open("postgres", "dbname=gorm")
	if err != nil {
		return nil, fmt.Errorf("unable to open db: %v", err)
	}

	return &storage{
		db: db,
	}, nil
}

// Insert provides inserting of data
func (s *storage) Insert(m *storage.LogRequest) error {
	return nil
}

// Close provides closing of db
func (s *storage) Close() error {
	return s.db.Close()
}
