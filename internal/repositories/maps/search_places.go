package maps

import (
	"context"

	"cloud.google.com/go/maps/places/apiv1/placespb"
	"github.com/dragonfish/go/v2/pkg/logger"
	"github.com/ride-app/catalog-service/internal/core/models"
	"google.golang.org/genproto/googleapis/type/latlng"
)

func (repo *Impl) ListPlaces(
	ctx context.Context,
	logger logger.Logger,
	searchString string,
	location *models.LatLng,
	sessionToken string,
) ([]models.Place, error) {
	req := &placespb.AutocompletePlacesRequest{
		Input:                   searchString,
		SessionToken:            sessionToken,
		IncludeQueryPredictions: false,
	}

	if location != nil {
		req.LocationBias = &placespb.AutocompletePlacesRequest_LocationBias{
			Type: &placespb.AutocompletePlacesRequest_LocationBias_Circle{
				Circle: &placespb.Circle{
					Center: &latlng.LatLng{
						Latitude:  location.Latitude,
						Longitude: location.Longitude,
					},
					Radius: 5000,
				},
			},
		}

		req.Origin = &latlng.LatLng{
			Latitude:  location.Latitude,
			Longitude: location.Longitude,
		}
	}

	res, err := repo.placesApi.AutocompletePlaces(ctx, req)
	if err != nil {
		return nil, err
	}

	places := make([]models.Place, len(res.Suggestions))

	for _, suggestion := range res.Suggestions {
		place := models.Place{}
		prediction := suggestion.GetPlacePrediction()

		place.Id = prediction.PlaceId
		place.DisplayName = prediction.StructuredFormat.MainText.Text
		place.Address = prediction.StructuredFormat.SecondaryText.Text

		places = append(places, place)
	}

	return places, nil
}
