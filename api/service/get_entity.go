package service

import (
	"context"

	"connectrpc.com/connect"
	pb "github.com/ride-app/entity-service/api/gen/ride/entity/v1alpha1"
)

func (service *EntityServiceServer) GetEntity(ctx context.Context,
	req *connect.Request[pb.GetEntityRequest]) (*connect.Response[pb.GetEntityResponse], error) {

	return connect.NewResponse(&pb.GetEntityResponse{}), nil
}
