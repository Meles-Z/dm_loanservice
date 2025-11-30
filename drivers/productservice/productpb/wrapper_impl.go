package productpb

import (
	"context"
	RpcClient "dm_loanservice/drivers/grpc"
	ctxDM "dm_loanservice/drivers/utils/context"
	"fmt"
	"time"

	productPb "github.com/brianjobling/dm_proto/generated/productservice/productpb"
	"github.com/spf13/viper"
)

// SetupWrapper initializes the product service wrapper
func SetupWrapper(conf *viper.Viper) *productServiceWrapper {
	w := &productServiceWrapper{}

	fmt.Printf("Connecting to product service at: %s\n", conf.GetString("productservice.address"))
	fmt.Printf("TLS enabled: %v\n", conf.GetBool("productservice.tls_enabled"))

	w.rpcCon = RpcClient.NewGrpcConnection(RpcClient.Options{
		PathCredential: conf.GetString("productservice.path_credential"),
		ServerName:     conf.GetString("productservice.servername"),
		Address:        conf.GetString("productservice.address"),
		Timeout:        time.Duration(conf.GetInt64("productservice.timeout")),
	})

	if w.rpcCon == nil {
		panic("RpcConnection is nil")
	}
	if w.rpcCon.Conn == nil {
		panic("gRPC connection is nil")
	}

	w.client = productPb.NewProductServiceClient(w.rpcCon.Conn)

	// Test connection
	state := w.rpcCon.Conn.GetState()
	fmt.Printf("Initial product service connection state: %v\n", state)

	return w
}

type productServiceWrapper struct {
	rpcCon *RpcClient.RpcConnection
	client productPb.ProductServiceClient
}

// ProductRead fetches product information by ID
func (s *productServiceWrapper) ProductRead(ctxSess *ctxDM.Context, id string) (*productPb.ProductReadRes, error) {
	grpcCtx := s.rpcCon.CreateContext(context.Background(), ctxSess)
	return s.client.ProductRead(grpcCtx, &productPb.ProductReadReq{
		Id: id,
	})
}

// Find products by customer ID
func (s *productServiceWrapper) ProductByCustomerID(ctxSess *ctxDM.Context, id string) (*productPb.FindProductsByCustomerIDRes, error) {
	fmt.Println("FindByProductID called successfully")
	grpcCtx := s.rpcCon.CreateContext(context.Background(), ctxSess)
	out, err := s.client.FindProductsByCustomerID(grpcCtx, &productPb.FindProductsByCustomerIDReq{
		CustomerID: id,
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("FindByProductID called successfully")
	return out, nil
}

func (s *productServiceWrapper) ProductOverviews(ctxSess *ctxDM.Context) (*productPb.ProductOverviewRes, error) {
	grpcCtx := s.rpcCon.CreateContext(context.Background(), ctxSess)
	return s.client.ProductOverview(grpcCtx, &productPb.ProductOverviewReq{})
}
