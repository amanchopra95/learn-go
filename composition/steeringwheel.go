package main

import "fmt"

type SteeringWheel struct{}

func (sw SteeringWheel) TurnRight() {
	fmt.Println("Turning Right...")
}

func (sw SteeringWheel) TurnLeft() {
	fmt.Println("Turning Left...")
}
