package rpcService

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	pb "github.com/Sona-28/netxd_customer"
	interfaces "github.com/Sona-28/netxd_dal/netxd_dal_interfaces"
	models "github.com/Sona-28/netxd_dal/netxd_dal_models"
	"sync"
)

type RPCServer struct {
	mu sync.Mutex
	pb.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.ICustomer
	Mcoll           *mongo.Collection
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pb.CustomerData) (*pb.CustomerResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	customer := &models.Customer{
		CustomerId: req.CustomerId,
		FirstName:  req.Firstname,
		LastName:   req.Lastname,
		BankId:     req.BankId,
		Balance:    req.Balance,
	}
	res, err := CustomerService.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}
	
	return &pb.CustomerResponse{
		CustomerId: res.CustomerId,
		CreatedAt:  res.CreatedAt,
	},nil
}

func (s *RPCServer) GetCustomer(ctx context.Context, req *pb.CustomerID) (*pb.CustomerData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	res, err := CustomerService.GetCustomer(req.CustomerId)
	if err != nil {
		return nil, err
	}
	customer := &pb.CustomerData{
		CustomerId: res.CustomerId,
		Firstname:  res.FirstName,
		Lastname:   res.LastName,
		BankId:     res.BankId,
		Balance:    res.Balance,
		CreatedAt:  res.CreatedAt,
		UpdatedAt:  res.UpdatedAt,
		IsActive:   res.IsActive,
	}
	return customer, nil
}

func (s *RPCServer) UpdateCustomer(ctx context.Context, req *pb.UpdateCustomerRequest) (*pb.CustomerResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	res, err := CustomerService.UpdateCustomer(&models.UpdateCustomerRequest{CustomerId:req.Id, Topic: req.Topic, NewValue: req.Newvalue})
	if err != nil {
		return nil, err
	}
	return &pb.CustomerResponse{
		CustomerId: res.CustomerId,
		CreatedAt:  res.CreatedAt,
	},nil
}

func (s *RPCServer) DeleteCustomer(ctx context.Context, req *pb.CustomerID) (*pb.CustomerResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	res, err := CustomerService.DeleteCustomer(req.CustomerId)
	if err != nil {
		return nil, err
	}
	return &pb.CustomerResponse{
		CustomerId: res.CustomerId,
		CreatedAt:  res.CreatedAt,
	}, nil
}
