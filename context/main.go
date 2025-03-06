package main

import "fmt"

func pass(arr *[]int) {
	fmt.Println(*arr)
	*arr = append(*arr, 5)
}

func main() {
	var arr = []int{2, 3, 4}
	pass(&arr)
	fmt.Println(arr)
}
