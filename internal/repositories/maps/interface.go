package maps

import (
	"context"

	"github.com/dragonfish/go/v2/pkg/logger"
	"github.com/ride-app/catalog-service/internal/core/models"
)

type MapsRepository interface {
	ListPlaces(
		ctx context.Context,
		logger logger.Logger,
		searchString string,
		location *models.LatLng,
		sessionToken string,
	) ([]models.Place, error)
}
