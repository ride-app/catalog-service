package apihandlers

import (
	"github.com/dragonfish/go/v2/pkg/logger"
	placesRepository "github.com/ride-app/catalog-service/internal/repositories/places"
)

type CatalogServiceServer struct {
	logger           logger.Logger
	placesRepository placesRepository.PlacesRepository
}

func New(
	logger logger.Logger,
	placesRepository placesRepository.PlacesRepository,
) *CatalogServiceServer {
	return &CatalogServiceServer{
		logger:           logger,
		placesRepository: placesRepository,
	}
}
