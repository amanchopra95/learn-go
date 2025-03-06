package main

import "fmt"

type Truck struct {
	Engine
	Transmission
	SteeringWheel
}

func (t Truck) FourWheelDrive() {
	fmt.Println("Four Wheel Drive")
}
