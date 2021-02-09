package application

import (
	"github.com/gorilla/mux"
	application "github.com/raphaelkieling/orders/application/controller"
	usecase "github.com/raphaelkieling/orders/application/useCase"
)

// CreateRouter Criar o conjunto de rotas
func CreateRouter() *mux.Router {
	orderController := application.NewOrderController(*usecase.NewGetOneOrderUseCase())

	r := mux.NewRouter()
	r.HandleFunc("/orders/{id}", orderController.GetOne)
	return r
}
