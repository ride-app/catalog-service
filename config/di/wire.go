//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/ride-app/entity-service/api/service"
	"github.com/ride-app/entity-service/config"
	entityrepository "github.com/ride-app/entity-service/repositories/entity"
	"github.com/ride-app/entity-service/utils/logger"
)

func InitializeService(logger logger.Logger, config *config.Config) (*service.EntityServiceServer, error) {
	panic(
		wire.Build(
			entityrepository.NewSomeEntityRepository,
			wire.Bind(
				new(entityrepository.EntityRepository),
				new(*entityrepository.SomeImpl),
			),
			service.New,
		),
	)
}
