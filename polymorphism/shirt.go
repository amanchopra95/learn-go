package main

type Shirt struct {
	ProductDetails
	Size  string
	Color string
}

func (s Shirt) calculate() int64 {
	clothingDiscount := float64(s.Price) * .20
	return s.Price - int64(clothingDiscount)
}
