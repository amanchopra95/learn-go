package main

import "fmt"

type purchasable interface {
	calculate() int64
}

var cart []purchasable

func addToCart(products ...purchasable) {
	cart = append(cart, products...)
}

func getCartTotal() int64 {
	var total int64
	for _, product := range cart {
		total += product.calculate()
	}
	return total
}

func main() {
	myShirt := Shirt{ProductDetails{Price: 5000, Brand: "Nike"}, "XL", "Blue"}
	myMonitor := Monitor{ProductDetails{Price: 5000, Brand: "Dell"}, "32 inch", "4k"}
	myWine := Wine{ProductDetails{Price: 5000, Brand: "SomeBrand"}, "2000", "Red"}

	addToCart(myShirt, myMonitor, myWine)
	fmt.Println(getCartTotal())
}
