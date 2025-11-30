package grpc

import (
	"context"
	"dm_loanservice/drivers/utils"
	"dm_loanservice/internal/service/domain/account"
	"time"

	pb "github.com/brianjobling/dm_proto/generated/accountservice/accountpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func encodeAccountAddRes(_ context.Context, r interface{}) (interface{}, error) {
	errorResp, ok := r.(*utils.ApplicationError)
	if ok {
		return nil, status.Errorf(codes.InvalidArgument, errorResp.Message)
	}
	acc, ok := r.(*account.AccountResponse)
	if !ok {
		return nil, status.Errorf(codes.Internal, "invalid response type")
	}
	return &pb.AccountCreateRes{
		Account: &pb.Account{
			Id:         acc.ID,
			CustomerId: acc.CustomerID,
			ProductId:  acc.ProductID,
			StartDate:  acc.StartDate.Format(time.RFC3339),
			EndDate:    acc.EndDate.Format(time.RFC3339),
		},
	}, nil
}

func encodeAccountReadRes(_ context.Context, r interface{}) (interface{}, error) {
	errorResp, ok := r.(*utils.ApplicationError)
	if ok {
		return nil, status.Errorf(codes.InvalidArgument, errorResp.Message)
	}
	acc := r.(*account.AccountReadResponse)
	return &pb.AccountGetRes{
		Account: &pb.Account{
			Id:         acc.Account.ID,
			CustomerId: acc.Account.CustomerID,
			ProductId:  acc.Account.ProductID,
			StartDate:  acc.Account.StartDate.Format(time.RFC3339),
			EndDate:    acc.Account.EndDate.Format(time.RFC3339),
			
		},
	}, nil
}
