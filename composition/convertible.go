package main

import "fmt"

type transmission interface {
	ShiftUp()
	ShiftDown()
}

type Convertible struct {
	Engine
	transmission
	SteeringWheel
}

func (c Convertible) ConvertTop() {
	fmt.Println("Convert Top")
}
