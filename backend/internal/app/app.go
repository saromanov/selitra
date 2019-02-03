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

	store, err := postgresql.Create()
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
	return nil
}

// GetLevelsStat returns map of levels for events
func (a *App) GetLevelsStat() sync.Map {
	return a.levelsStat
}
