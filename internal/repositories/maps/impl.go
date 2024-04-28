package maps

import (
	places "cloud.google.com/go/maps/places/apiv1"
	"github.com/dragonfish/go/v2/pkg/logger"
	"github.com/ride-app/catalog-service/config"
)

type Impl struct {
	placesApi *places.Client
}

func New(log logger.Logger, config *config.Config, placesApi *places.Client) (*Impl, error) {
	defer log.Info("Places Repository initialized")
	return &Impl{placesApi: placesApi}, nil
}
