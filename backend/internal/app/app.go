package app

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/saromanov/selitra/backend/internal/storage"
	"github.com/saromanov/selitra/backend/internal/storage/postgresql"
	structs "github.com/saromanov/selitra/backend/internal/structs/v1"
)

// App defines main logic
type App struct {
	mu          *sync.RWMutex
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
		mu: &sync.RWMutex{},
	}, nil
}

// SendEvent provides sending of the event
func (a *App) SendEvent(r *structs.LogRequest) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	atomic.AddUint32(&a.eventsCount, 1)
	if err := a.db.Insert(logRequestToModel(r)); err != nil {
		return fmt.Errorf("unable to insert data: %v", err)
	}
	return nil
}

// Search returns list of events
func (a *App) Search(s *structs.SearchRequest) ([]*structs.LogRequest, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	result, err := a.db.Search(searchRequestToInner(s))
	if err != nil {
		return nil, err
	}
	return searchModelsToResponse(result), nil
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

func searchModelsToResponse(s []*storage.LogRequest) []*structs.LogRequest {
	result := make([]*structs.LogRequest, len(s))
	for i, item := range s {
		result[i] = modelToLogRequest(item)
	}
	return result
}

func modelToLogRequest(r *storage.LogRequest) *structs.LogRequest {
	return &structs.LogRequest{
		Message:   r.Message,
		Name:      r.Name,
		Timestamp: int64(r.Timestamp),
		Entry:     r.Entry,
		Service:   r.Service,
		Labels:    r.Labels,
	}
}

func searchRequestToInner(r *structs.SearchRequest) *storage.SearchRequest {
	resp := &storage.SearchRequest{}
	if r.FromTimestamp != "" && r.ToTimestamp != "" {
		resp.FromTimestamp, _ = strconv.ParseInt(r.FromTimestamp, 0, 64)
		resp.ToTimestamp, _ = strconv.ParseInt(r.ToTimestamp, 0, 64)
	}
	resp.Name = r.Name
	return resp
}
