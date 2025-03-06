package main

type Monitor struct {
	ProductDetails
	Size  string
	Color string
}

func (s Monitor) calculate() int64 {
	electronicTax := float64(s.Price) * .30
	return s.Price + int64(electronicTax)
}
