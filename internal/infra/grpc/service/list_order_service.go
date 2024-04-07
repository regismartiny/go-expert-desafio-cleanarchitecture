package service

import (
	"context"

	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type ListOrdersService struct {
	pb.UnimplementedListOrdersServiceServer
	ListOrderUseCase usecase.ListOrdersUseCase
}

func NewListOrdersService(listOrdersUseCase usecase.ListOrdersUseCase) *ListOrdersService {
	return &ListOrdersService{
		ListOrderUseCase: listOrdersUseCase,
	}
}

func (s *ListOrdersService) ListOrders(ctx context.Context, in *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	output, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	orders := []*pb.Order{}

	for _, order := range output.Orders {
		orders = append(orders, &pb.Order{
			Id:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return &pb.ListOrdersResponse{
		Quantity: int32(output.Quantity),
		Orders:   orders,
	}, nil
}
