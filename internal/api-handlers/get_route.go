package apihandlers

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	pb "github.com/ride-app/catalog-service/api/ride/catalog/v1alpha1"
)

func (service *CatalogServiceServer) GetRoute(ctx context.Context,
	req *connect.Request[pb.GetRouteRequest],
) (*connect.Response[pb.GetRouteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
}
