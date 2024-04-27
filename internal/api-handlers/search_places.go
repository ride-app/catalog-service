package apihandlers

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/aidarkhanov/nanoid"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/catalog-service/api/ride/catalog/v1alpha1"
	"github.com/ride-app/catalog-service/internal/core/models"
	"google.golang.org/genproto/googleapis/type/latlng"
)

func (service *CatalogServiceServer) SearchPlaces(ctx context.Context,
	req *connect.Request[pb.SearchPlacesRequest],
) (*connect.Response[pb.SearchPlacesResponse], error) {
	log := service.logger.WithFields(map[string]string{
		"method": "SearchPlaces",
	})

	validator, err := protovalidate.New()
	if err != nil {
		log.WithError(err).Info("Failed to initialize validator")

		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if err := validator.Validate(req.Msg); err != nil {
		log.WithError(err).Info("Invalid request")

		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	var sessionToken string

	if req.Msg.SessionToken == nil {
		sessionToken = nanoid.New()
	} else {
		sessionToken = *req.Msg.SessionToken
	}

	places, err := service.placesRepository.ListPlaces(
		ctx,
		log,
		req.Msg.Query,
		models.LatLngFromProtobuf(req.Msg.Location),
		sessionToken,
	)
	if err != nil {
		log.WithError(err).Error("failed to get places")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	placeProtobufs := make([]*pb.Place, len(places))

	for _, place := range places {
		protobufPlace := pb.Place{}

		protobufPlace.DisplayName = place.DisplayName
		protobufPlace.Address = place.Address
		protobufPlace.Location = &latlng.LatLng{
			Latitude:  place.Location.Latitude,
			Longitude: protobufPlace.Location.Longitude,
		}
		protobufPlace.DistanceMeters = place.DistanceMeters
		protobufPlace.Name = fmt.Sprintf("places/%s", place.Id)

		switch place.Type {
		case models.PLACE_TYPE_UNSPECIFIED:
			protobufPlace.Type = pb.Place_TYPE_UNSPECIFIED
		case models.PLACE_TYPE_PICKUP:
			protobufPlace.Type = pb.Place_TYPE_PICKUP
		case models.PLACE_TYPE_DROPOFF:
			protobufPlace.Type = pb.Place_TYPE_DROPOFF
		}

		placeProtobufs = append(placeProtobufs, &protobufPlace)
	}

	res := connect.NewResponse(&pb.SearchPlacesResponse{
		SessionToken: sessionToken,
		Places:       placeProtobufs,
	})

	if err := validator.Validate(res.Msg); err != nil {
		log.WithError(err).Error("Invalid response")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return res, nil
}
