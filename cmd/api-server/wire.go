//go:build wireinject

package main

import (
	"github.com/dragonfish/go/v2/pkg/logger"
	"github.com/google/wire"
	"github.com/ride-app/catalog-service/config"
	apihandlers "github.com/ride-app/catalog-service/internal/api-handlers"
	placesrepository "github.com/ride-app/catalog-service/internal/repositories/places"
	thirdparty "github.com/ride-app/catalog-service/third-party"
)

func InitializeService(
	logger logger.Logger,
	config *config.Config,
) (*apihandlers.CatalogServiceServer, error) {
	panic(
		wire.Build(
			thirdparty.NewPlacesClient,
			placesrepository.New,
			wire.Bind(
				new(placesrepository.PlacesRepository),
				new(*placesrepository.Impl),
			),
			apihandlers.New,
		),
	)
}
