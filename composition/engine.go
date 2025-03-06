package main

import "fmt"

type Engine struct {
}

func (e Engine) Start() {
	fmt.Println("Engine Started")
}
