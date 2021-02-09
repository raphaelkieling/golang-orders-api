package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	usecase "github.com/raphaelkieling/orders/application/useCase"
)

// OrderController controller para pedidos
type OrderController struct {
	useCase usecase.GetOneOrderUseCase
}

// GetOne pega um pedido
func (o *OrderController) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	order := o.useCase.Execute(vars["id"])
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(order.ExternalID))
}

// NewOrderController cria uma controller
func NewOrderController(getOneUseCase usecase.GetOneOrderUseCase) *OrderController {
	return &OrderController{
		useCase: getOneUseCase,
	}
}
