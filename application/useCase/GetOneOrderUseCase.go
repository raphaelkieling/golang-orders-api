package usecase

import (
	domain "github.com/raphaelkieling/orders/domain"
)

// GetOneOrderUseCase controller para pedidos
type GetOneOrderUseCase struct {
}

// Execute pega um pedido
func (g *GetOneOrderUseCase) Execute(orderID string) *domain.Order {
	if orderID == "2" {
		return &domain.Order{
			ExternalID: "xxx",
		}
	}
	return &domain.Order{
		ExternalID: "0002020392",
	}
}

// NewGetOneOrderUseCase criar um use case
func NewGetOneOrderUseCase() *GetOneOrderUseCase {
	return &GetOneOrderUseCase{}
}
