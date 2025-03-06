package main

import (
	"fmt"
	"sync"
)

func someFunc(myChan chan int, wg *sync.WaitGroup, num int) {
	defer wg.Done()
	fmt.Println(num, "Minus 1")
	myChan <- num
}

func closeChan(ch chan int, wg *sync.WaitGroup) {
	wg.Wait()
	close(ch)
}

func join() {
	myChan := make(chan int)
	arr := []int{}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go someFunc(myChan, wg, 2)
	wg.Add(1)
	go someFunc(myChan, wg, 3)
	wg.Add(1)
	go someFunc(myChan, wg, 4)
	fmt.Println("Called goroutines and now waiting")

	go closeChan(myChan, wg)

	for n := range myChan {
		arr = append(arr, n)
	}

	fmt.Println("Done", arr)
}

func asyncComm() {
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, s := range chars {
			select {
			case charChannel <- s:
			}
		}
	}()

	wg.Wait()
	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}
}
