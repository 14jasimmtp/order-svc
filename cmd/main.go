package main

import (
	"fmt"
	"log"
	"net"

	"github.com/14jasimmtp/order-svc/pkg/client"
	"github.com/14jasimmtp/order-svc/pkg/config"
	"github.com/14jasimmtp/order-svc/pkg/db"
	"github.com/14jasimmtp/order-svc/pkg/pb"
	order "github.com/14jasimmtp/order-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal("error while connecting env ", err)
	}

	h := db.Connection(c.DB_URL)

	lis, err := net.Listen("tcp", c.PORT)
	if err != nil {
		log.Fatalln("Failed to listen : ", err)
	}

	ProductService := client.InitProductServiceClient(c.PRODUCT_SVC_URL)
	fmt.Println("Order service on ", c.PORT)

	s := order.Server{
		H:          h,
		ProductSvc: ProductService,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer,&s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve : ", err)
	}

}
