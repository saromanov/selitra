package app

import (
	"sync"
	"sync/atomic"
)

// App defines main logic
type App struct {
	mu          *sync.RWMutex
	levelsStat  sync.Map
	eventsCount uint32
}

// New provides initialization of the app
func New() *App {
	return &App{}
}

// SendEvent provides sending of the event
func (a *App) SendEvent(r *LogRequest) error {
	a.levelsStat.Store(r.Level, 0)
	atomic.AddUint32(&a.eventsCount, 1)
	return nil
}

// GetLevelsStat returns map of levels for events
func (a *App) GetLevelsStat() sync.Map {
	return a.levelsStat
}
