package domain

type Product struct {
	ID       string
	Name     string
	Price    float64
	OrderID  string
	Quantity int
}

func (p *Product) Total() float64 {
	if p.OrderID != "" {
		return p.Price * float64(p.Quantity)
	}
	return p.Price
}
