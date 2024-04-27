package apihandlers

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	pb "github.com/ride-app/catalog-service/api/ride/catalog/v1alpha1"
)

func (service *CatalogServiceServer) GetCatalog(ctx context.Context,
	req *connect.Request[pb.GetCatalogRequest],
) (*connect.Response[pb.GetCatalogResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
}
