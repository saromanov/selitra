package postgresql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// storage implements db handling with Postgesql
type storage struct {
	db *gorm.DB
}

func Create() (*storage, error) {
	db, err := gorm.Open("postgres", "dbname=gorm")
	if err != nil {
		return nil, fmt.Errorf("unable to open db: %v", err)
	}

	return &storage{
		db: db,
	}, nil
}

func (s *storage) Close() {
	s.db.Close()
}
