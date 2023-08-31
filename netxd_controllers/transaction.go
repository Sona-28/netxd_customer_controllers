package rpcService

import (
	"context"
	interfaces "github.com/Sona-28/netxd_dal/netxd_dal_interfaces"
	pb "github.com/Sona-28/netxd_transaction"
	"sync"
)

type TransactionSever struct {
	mu sync.Mutex
	pb.UnimplementedTransactionServiceServer
}

var (
	TransactionService interfaces.TransactionInterface
)

func (t *TransactionSever) TransferMoney(ctx context.Context, req *pb.TransactionData) (*pb.TransactionResponse, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	res, err := TransactionService.TransferMoney(req.From, req.To, req.Amount)
	if err != nil {
		return nil, err
	}
	return &pb.TransactionResponse{
		Message: res,
	}, nil
}
