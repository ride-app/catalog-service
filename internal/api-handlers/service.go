package apihandlers

import (
	"github.com/dragonfish/go/v2/pkg/logger"
	mapsRepository "github.com/ride-app/catalog-service/internal/repositories/maps"
)

type CatalogServiceServer struct {
	logger         logger.Logger
	mapsRepository mapsRepository.MapsRepository
}

func New(
	logger logger.Logger,
	mapsRepository mapsRepository.MapsRepository,
) *CatalogServiceServer {
	return &CatalogServiceServer{
		logger:         logger,
		mapsRepository: mapsRepository,
	}
}
