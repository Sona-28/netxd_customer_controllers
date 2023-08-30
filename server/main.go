package main

import (
	"context"
	"fmt"
	"net"
	"netxd_project/netxd_customer_controllers/config"
	"netxd_project/netxd_customer_controllers/constants"
	rpc "netxd_project/netxd_customer_controllers/netxd_controllers"
	pb "netxd_project/netxd_customer"
	service "netxd_project/netxd_dal/netxd_dal_services"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)


func initApp(mongoClient *mongo.Client){
	rpc.Mcoll = config.GetCollection(mongoClient, constants.Dbname, "customer")
	rpc.CustomerService = service.InitCustomer(rpc.Mcoll, context.Background())
}

func main() {
	mongoClient,err := config.ConnectDataBase()
	defer mongoClient.Disconnect(context.TODO())
	if err!=nil{
		panic(err)
	}
	initApp(mongoClient)
	lis, err := net.Listen("tcp", constants.Port)
	fmt.Println("Server listening on: ", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen:%v", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterCustomerServiceServer(s,&rpc.RPCServer{})
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve:%v", err)
	}
	fmt.Println("finish")
	
	
}