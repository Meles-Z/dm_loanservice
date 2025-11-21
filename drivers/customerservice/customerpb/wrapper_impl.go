package customerpb

import (
	"context"
	RpcClient "dm_loanservice/drivers/grpc"
	ctxDM "dm_loanservice/drivers/utils/context"
	"fmt"
	"time"

	"github.com/brianjobling/dm_proto/generated/userservice/customerpb"
	"github.com/spf13/viper"
)

// In dm_loanservice/drivers/customerservice/customerpb/wrapper_impl.go
func SetupWrapper(conf *viper.Viper) *customerServiceWrapper {
	w := &customerServiceWrapper{}

	fmt.Printf("Connecting to customer service at: %s\n", conf.GetString("customerservice.address"))
	fmt.Printf("TLS enabled: %v\n", conf.GetBool("customerservice.tls_enabled"))

	w.rpcCon = RpcClient.NewGrpcConnection(RpcClient.Options{
		PathCredential: conf.GetString("customerservice.path_credential"),
		ServerName:     conf.GetString("customerservice.servername"),
		Address:        conf.GetString("customerservice.address"),
		Timeout:        time.Duration(conf.GetInt64("customerservice.timeout")),
	})

	if w.rpcCon == nil {
		panic("RpcConnection is nil")
	}
	if w.rpcCon.Conn == nil {
		panic("gRPC connection is nil")
	}

	w.client = customerpb.NewCustomerServiceClient(w.rpcCon.Conn)

	// Test the connection state
	state := w.rpcCon.Conn.GetState()
	fmt.Printf("Initial connection state: %v\n", state)

	return w
}

type customerServiceWrapper struct {
	rpcCon *RpcClient.RpcConnection
	client customerpb.CustomerServiceClient
	config *viper.Viper
}

func (s *customerServiceWrapper) FindByProductID(ctxSess *ctxDM.Context, productID string) (resp []*customerpb.Customers, err error) {
	fmt.Println("Here is the first pralace of grpc")
	grpcCtx := s.rpcCon.CreateContext(context.Background(), ctxSess)
	out, err := s.client.CustomerByProductID(grpcCtx, &customerpb.CustomerByProductReq{
		ProductID: productID,
	})
	if err != nil {
		return nil, err
	}
	resp = out.Customers
	fmt.Println("Called callled callled::::")
	return
}
