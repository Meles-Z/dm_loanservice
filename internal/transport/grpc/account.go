package grpc

import (
	"context"
	"dm_loanservice/internal/endpoint"

	"github.com/brianjobling/dm_proto/generated/accountservice/accountpb"
	gt "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type accountServiceServer struct {
	add                 gt.Handler
	get                 gt.Handler
	accountArreas       gt.Handler
	accountEdit         gt.Handler
	accountDelete       gt.Handler
	accountList         gt.Handler
	recentArrearsCases  gt.Handler
	recentArrears       gt.Handler
	mortgagePerformance gt.Handler
}

func NewAccountServer(endpoints endpoint.Endpoints) accountpb.AccountServiceServer {
	return &accountServiceServer{
		add:           gt.NewServer(endpoints.AccountAdd, decodeAccountAddReq, encodeAccountAddRes),
		get:           gt.NewServer(endpoints.AccountRead, decodeAccountReadReq, encodeAccountReadRes),
		accountArreas: gt.NewServer(nil, nil, nil),
	}
}

func accountDeadlines(ctx context.Context) error {
	if ctx.Err() == context.DeadlineExceeded {
		return status.Error(codes.Canceled, "The client canceled the request!")
	} else {
		return nil
	}
}

func (s *accountServiceServer) CreateAccount(ctx context.Context, req *accountpb.AccountCreateReq) (*accountpb.AccountCreateRes, error) {
	errDeadline := accountDeadlines(ctx)
	if errDeadline != nil {
		return nil, errDeadline
	}
	_, res, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*accountpb.AccountCreateRes), nil
}

func (s *accountServiceServer) GetAccount(ctx context.Context, req *accountpb.AccountGetReq) (*accountpb.AccountGetRes, error) {
	errDeadline := accountDeadlines(ctx)
	if errDeadline != nil {
		return nil, errDeadline
	}
	_, res, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*accountpb.AccountGetRes), nil
}

func (s *accountServiceServer) AccountArrears(ctx context.Context, req *accountpb.AccountArrearsReq) (*accountpb.AccountArrearsRes, error) {
	errDeadline := accountDeadlines(ctx)
	if errDeadline != nil {
		return nil, errDeadline
	}
	_, res, err := s.accountArreas.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*accountpb.AccountArrearsRes), nil
}

func (s *accountServiceServer) UpdateAccount(ctx context.Context, req *accountpb.AccountUpdateReq) (*accountpb.AccountUpdateRes, error) {
	errDeadline := accountDeadlines(ctx)
	if errDeadline != nil {
		return nil, errDeadline
	}
	_, res, err := s.accountEdit.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*accountpb.AccountUpdateRes), nil
}

func (s *accountServiceServer) DeleteAccount(ctx context.Context, req *accountpb.AccountDeleteReq) (*accountpb.AccountDeleteRes, error) {
	errDeadline := accountDeadlines(ctx)
	if errDeadline != nil {
		return nil, errDeadline
	}
	_, res, err := s.accountDelete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*accountpb.AccountDeleteRes), nil
}

func (s *accountServiceServer) ListAccounts(ctx context.Context, req *accountpb.AccountListReq) (*accountpb.AccountListRes, error) {
	errDeadline := accountDeadlines(ctx)
	if errDeadline != nil {
		return nil, errDeadline
	}
	_, res, err := s.accountList.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*accountpb.AccountListRes), nil

}

func (s *accountServiceServer) RecentArrearsCases(ctx context.Context, req *accountpb.RecentArrearsCasesReq) (*accountpb.RecentArrearsCasesRes, error) {
	errDeadline := accountDeadlines(ctx)
	if errDeadline != nil {
		return nil, errDeadline
	}
	_, res, err := s.recentArrearsCases.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*accountpb.RecentArrearsCasesRes), nil
}

func (s *accountServiceServer) RecentArrears(ctx context.Context, req *accountpb.RecentArrearsReq) (*accountpb.RecentArrearsRes, error) {
	errDeadline := accountDeadlines(ctx)
	if errDeadline != nil {
		return nil, errDeadline
	}
	_, res, err := s.recentArrears.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*accountpb.RecentArrearsRes), nil
}

func (s *accountServiceServer) MortgagePerformance(ctx context.Context, req *accountpb.MortgagePerformanceReq) (*accountpb.MortgagePerformanceRes, error) {
	errDeadline := accountDeadlines(ctx)
	if errDeadline != nil {
		return nil, errDeadline
	}
	_, res, err := s.mortgagePerformance.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*accountpb.MortgagePerformanceRes), nil
}
