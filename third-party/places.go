package thirdparty

import (
	"context"

	places "cloud.google.com/go/maps/places/apiv1"
	"github.com/dragonfish/go/v2/pkg/logger"
	"github.com/ride-app/catalog-service/config"
	"google.golang.org/api/option"
)

func NewPlacesClient(log logger.Logger, config *config.Config) (*places.Client, error) {
	ctx := context.Background()
	client, err := places.NewClient(ctx, option.WithAPIKey(config.MapsApiKey))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}
