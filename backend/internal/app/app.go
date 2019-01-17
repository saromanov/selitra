package app

// App defines main logic
type App struct {
	eventsCount map[string]uint64
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
