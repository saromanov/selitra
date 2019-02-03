package app

import (
	"sync"
	"sync/atomic"

	structs "github.com/saromanov/selitra/backend/internal/structs/v1"
	"github.com/saromanov/selitra/backend/internal/storage"
	"github.com/saromanov/selitra/backend/internal/storage/postgres"
)

// App defines main logic
type App struct {
	mu          *sync.RWMutex
	levelsStat  sync.Map
	eventsCount uint32
	db storage.Storage
}

// New provides initialization of the app
func New() (*App, error) {

	store, err := postgres.Create()
	if err != nil {
		return nil, fmt.Errorf("unable to setup Postgresql: %v", err)
	}
	return &App{
		db: store,
	}
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
