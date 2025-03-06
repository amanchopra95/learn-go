package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func doWork() int {
	time.Sleep(time.Second)
	return rand.IntN(100)
}

func main() {

	dataChan := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := doWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Println(n)
	}

}
