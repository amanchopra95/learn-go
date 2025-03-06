package main

type Wine struct {
	ProductDetails
	Size  string
	Color string
}

func (s Wine) calculate() int64 {
	liquorTax := float64(s.Price) * .25
	stateTax := float64(s.Price) * .10
	return s.Price + int64(liquorTax) + int64(stateTax)
}
