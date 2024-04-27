package service

import (
	"context"

	"connectrpc.com/connect"
	pb "github.com/ride-app/entity-service/api/gen/ride/entity/v1alpha1"
)

func (service *EntityServiceServer) ListEntities(ctx context.Context,
	req *connect.Request[pb.ListEntitiesRequest]) (*connect.Response[pb.ListEntitiesResponse], error) {

	return connect.NewResponse(&pb.ListEntitiesResponse{}), nil
}
