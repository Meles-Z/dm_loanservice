package grpc

import (
	"context"

	"github.com/brianjobling/dm_proto/generated/customerservice/customerpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type CustomerClient interface {
	FindByProductID(ctx context.Context, id string, token string) ([]*customerpb.Customers, error)
}

type customerClient struct {
	c customerpb.CustomerServiceClient
}

func NewCustomerClient(conn *grpc.ClientConn) CustomerClient {
	return &customerClient{
		c: customerpb.NewCustomerServiceClient(conn),
	}
}

func (cl *customerClient) FindByProductID(ctx context.Context, id string, token string) ([]*customerpb.Customers, error) {

	// Add Bearer token metadata
	md := metadata.New(map[string]string{
		"authorization": "Bearer " + token,
	})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := cl.c.CustomerByProductID(ctx, &customerpb.CustomerByProductReq{
		ProductID: id,
	})
	if err != nil {
		return nil, err
	}

	return resp.Customers, nil
}
