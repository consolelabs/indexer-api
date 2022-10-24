package handler

import (
	"github.com/consolelabs/indexer-api/pkg/entity"
	"github.com/consolelabs/indexer-api/pkg/logger"
	"github.com/consolelabs/indexer-api/pkg/queue"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
)

type Handler struct {
	store    *store.Store
	service  *service.Service
	logger   logger.Logger
	entities *entity.Entity
	queue    *queue.Queue
}

func New(store *store.Store, service *service.Service, queue *queue.Queue) (*Handler, error) {
	return &Handler{
		store:    store,
		service:  service,
		logger:   logger.L,
		entities: entity.New(store, service),
		queue:    queue,
	}, nil
}
