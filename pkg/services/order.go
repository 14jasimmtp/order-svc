package order

import (
	"context"
	"fmt"
	"net/http"

	"github.com/14jasimmtp/order-svc/pkg/client"
	"github.com/14jasimmtp/order-svc/pkg/db"
	"github.com/14jasimmtp/order-svc/pkg/models"
	"github.com/14jasimmtp/order-svc/pkg/pb"
)

type Server struct {
	H          db.Handler
	ProductSvc client.ProductServiceClient
	pb.UnimplementedOrderServiceServer
}

func (s *Server) CreateOrder(c context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	product, err := s.ProductSvc.FindOne(req.ProductId)
	fmt.Println(product.Data,req.Quantity)
	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	} else if product.Data.Price < req.Quantity {
		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: "Stock too less"}, nil
	}

	order := models.Order{
		Price:     product.Data.Price,
		ProductId: product.Data.Id,
		UserId:    req.UserId,
	}

	s.H.DB.Create(&order)

	res, err := s.ProductSvc.DecreaseStock(req.ProductId, order.Id)

	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	} else if res.Status == http.StatusConflict {
		s.H.DB.Delete(&models.Order{}, order.Id)

		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: res.Error}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}
