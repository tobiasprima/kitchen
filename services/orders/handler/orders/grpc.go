package handler

import (
	"context"

	"github.com/tobiasprima/kitchen/services/common/genproto/orders"
	"github.com/tobiasprima/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	ordersService types.OrderService

	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server ,ordersService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		ordersService: ordersService,
	}

	// register OrderServiceServer
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)


}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderId: 42,
		CustomerID: 2,
		ProductID: 1,
		Quantity: 10,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}