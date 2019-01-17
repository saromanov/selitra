package app

import (
	log "github.com/Sirupsen/logrus"
)

// App defines main logic
type App struct {
	eventsCount map[string]uint64
	levels      []log.Level
}

// New provides initialization of the app
func New() *App {
	return &App{
		eventsCount: make(map[string]uint64),
	}
}

// GetEventsCount returns map of counts for events
func (a *App) GetEventsCount() map[string]uint64 {
	return a.eventsCount
}

// Levels retruns current usage levels
func (a *App) Levels() []log.Level {
	return a.levels
}

// SetLevels provides setting of leveles to teh app
func (a *App) SetLevels(levels []log.Level) {
	a.levels = levels
}

// Fire method implemented by Logrus Hook interface
func (a *App) Fire(e *log.Entry) error {

	return nil
}
