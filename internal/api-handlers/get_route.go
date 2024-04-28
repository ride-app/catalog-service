package apihandlers

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/ride-app/catalog-service/api/ride/catalog/v1alpha1"
)

func (service *CatalogServiceServer) GetRoute(ctx context.Context,
	req *connect.Request[pb.GetRouteRequest],
) (*connect.Response[pb.GetRouteResponse], error) {
	log := service.logger.WithFields(map[string]string{
		"method": "GetRoute",
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

	// switch x := req.Msg.Origin.(type) {
	// case *pb.GetRouteRequest_Location_Place:

	// }

	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
}
