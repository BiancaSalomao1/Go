package main

import (
	"fmt"
	"pedidos/internal/domain"
)

func main() {
	product := domain.Product{
		ID:       "1",
		Name:     "Product 1",
		Price:    100.0,
		Quantity: 1,
	}

	order := domain.Order{
		ID:          "1",
		CustomerID:  "1",
		Products:    []domain.Product{product},
		OrderStatus: domain.OrderStatusPending,
	}

	order.AddProduct(product)

	fmt.Printf("Pedido: %+v\n", order)
	fmt.Printf("Total: R$ %.2f\n", order.Total())
}
