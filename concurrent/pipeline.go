package main

import (
	"fmt"
)

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func squareTheNumber(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func pipeline() {
	nums := []int{2, 3, 4, 5, 6, 7}

	// Stage 1
	// Slice to channel
	dataChannel := sliceToChannel(nums)
	// Stage 2
	// Square each number
	squareChannel := squareTheNumber(dataChannel)
	// Stage 3
	// Display result
	for n := range squareChannel {
		fmt.Println(n)
	}
}
