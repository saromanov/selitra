package postgresql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	st "github.com/saromanov/selitra/backend/internal/storage"
	structs "github.com/saromanov/selitra/backend/internal/structs/v1"
)

// storage implements db handling with Postgesql
type storage struct {
	db *gorm.DB
}

// Create provides init for postgesql storage
func Create(s *structs.Config) (st.Storage, error) {
	args := "dbname=selitra"
	if s.DBName != "" && s.DBPassword != "" && s.DBUser != "" {
		args += fmt.Sprintf(" user=%s dbname=%s password=%s", s.DBUser, s.DBName, s.DBPassword)
	}
	db, err := gorm.Open("postgres", args)
	if err != nil {
		return nil, fmt.Errorf("unable to open db: %v", err)
	}
	db.AutoMigrate(&st.LogRequest{})
	return &storage{
		db: db,
	}, nil
}

// Insert provides inserting of data
func (s *storage) Insert(m *st.LogRequest) error {
	err := s.db.Create(m).Error
	if err != nil {
		return fmt.Errorf("storage: unable to insert data: %v", err)
	}
	return nil
}

// Insert provides finding data
func (s *storage) Search(sr *st.SearchRequest) ([]*st.LogRequest, error) {
	var response []*st.LogRequest
	err := s.makeQuery(s.db, sr).Find(&response).Error
	if err != nil {
		return nil, fmt.Errorf("storage: unable to find data: %v", err)
	}
	return response, nil
}

func (s *storage) makeQuery(db *gorm.DB, sr *st.SearchRequest) *gorm.DB {
	if sr.Name != "" {
		db = db.Where("name=?", sr.Name)
		return db
	}
	return db
}

// Aggregate provides aggregation on the data
func (s *storage) Aggregate(sr *st.SearchRequest) ([]*st.LogRequest, error) {
	return nil, nil
}

// Close provides closing of db
func (s *storage) Close() error {
	return s.db.Close()
}
