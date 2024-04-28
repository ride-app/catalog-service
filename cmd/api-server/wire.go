//go:build wireinject

package main

import (
	"github.com/dragonfish/go/v2/pkg/logger"
	"github.com/google/wire"
	"github.com/ride-app/catalog-service/config"
	apihandlers "github.com/ride-app/catalog-service/internal/api-handlers"
	mapsrepository "github.com/ride-app/catalog-service/internal/repositories/maps"
	thirdparty "github.com/ride-app/catalog-service/third-party"
)

func InitializeService(
	logger logger.Logger,
	config *config.Config,
) (*apihandlers.CatalogServiceServer, error) {
	panic(
		wire.Build(
			thirdparty.NewPlacesClient,
			mapsrepository.New,
			wire.Bind(
				new(mapsrepository.MapsRepository),
				new(*mapsrepository.Impl),
			),
			apihandlers.New,
		),
	)
}
