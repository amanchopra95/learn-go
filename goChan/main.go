package main

import (
	"fmt"
	"time"
)

func timesThree(arr []int, ch chan int) {
	minusCh := make(chan int, 3)
	for _, elem := range arr {
		value := elem * 3
		if value%2 == 0 {
			go minusThree(value, minusCh)
			value = <-minusCh
		}
		ch <- value
	}
}

func minusThree(number int, ch chan int) {
	ch <- number - 3
	fmt.Println("The function continues after returning the result")
}

// var n = 1
// var mu sync.Mutex

// func timesThree() {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	n *= 3
// 	fmt.Println(n)
// }

func main() {
	fmt.Println("We are going to execute a goroutine")
	arr := []int{2, 3, 4}
	ch := make(chan int, len(arr)) // buffered channel:- a channel that can store value
	go timesThree(arr, ch)
	// for i := 0; i < 10; i++ {
	// 	// fmt.Printf("The result is: %v \n", <-ch)
	// 	go timesThree()
	// }
	time.Sleep(time.Second)
	fmt.Println("The result before goroutine")
}
