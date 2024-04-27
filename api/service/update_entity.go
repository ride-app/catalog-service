package service

import (
	"context"

	"connectrpc.com/connect"
	pb "github.com/ride-app/entity-service/api/gen/ride/entity/v1alpha1"
)

func (service *EntityServiceServer) UpdateEntity(ctx context.Context,
	req *connect.Request[pb.UpdateEntityRequest]) (*connect.Response[pb.UpdateEntityResponse], error) {

	return connect.NewResponse(&pb.UpdateEntityResponse{}), nil
}
