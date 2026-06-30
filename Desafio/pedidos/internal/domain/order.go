package domain

type Order struct {
	ID          string
	CustomerID  string
	Products    []Product
	OrderStatus OrderStatus
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusPaid      OrderStatus = "PAID"
	OrderStatusCancelled OrderStatus = "CANCELLED"
)

func (o *Order) AddProduct(product Product) {
	o.Products = append(o.Products, product)
}

func (o *Order) Total() float64 {
	total := 0.0
	for _, product := range o.Products {
		total += product.Price * float64(product.Quantity)
	}
	return total
}

func (s OrderStatus) OrderStatus() string {
	switch s {
	case OrderStatusPending:
		return "PENDING"
	case OrderStatusPaid:
		return "PAID"
	case OrderStatusCancelled:
		return "CANCELLED"
	default:
		return "UNKNOWN"
	}
}
