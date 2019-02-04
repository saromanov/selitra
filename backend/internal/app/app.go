package app

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/saromanov/selitra/backend/internal/storage"
	"github.com/saromanov/selitra/backend/internal/storage/postgresql"
	structs "github.com/saromanov/selitra/backend/internal/structs/v1"
)

// App defines main logic
type App struct {
	mu          *sync.RWMutex
	levelsStat  sync.Map
	eventsCount uint32
	db          storage.Storage
}

// New provides initialization of the app
func New(c *structs.Config) (*App, error) {

	store, err := postgresql.Create(c)
	if err != nil {
		return nil, fmt.Errorf("unable to setup Postgresql: %v", err)
	}
	return &App{
		db: store,
	}, nil
}

// SendEvent provides sending of the event
func (a *App) SendEvent(r *structs.LogRequest) error {
	a.levelsStat.Store(r.Level, 0)
	atomic.AddUint32(&a.eventsCount, 1)
	if err := a.db.Insert(logRequestToModel(r)); err != nil {
		return fmt.Errorf("unable to insert data: %v", err)
	}
	return nil
}

// GetEvents returns list of events
func (a *App) GetEvents() ([]*structs.LogRequest, error) {
	return nil, nil
}

// GetLevelsStat returns map of levels for events
func (a *App) GetLevelsStat() sync.Map {
	return a.levelsStat
}

func logRequestToModel(r *structs.LogRequest) *storage.LogRequest {
	return &storage.LogRequest{
		Message:   r.Message,
		Name:      r.Name,
		Timestamp: uint64(r.Timestamp),
		Entry:     r.Entry,
		Service:   r.Service,
		Labels:    r.Labels,
	}
}
