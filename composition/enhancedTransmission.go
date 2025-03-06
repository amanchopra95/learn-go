package main

import "fmt"

type EnhancedTransmission struct{}

func (et EnhancedTransmission) ShiftUp() {
	fmt.Println("Enhanced Shift UP")
}

func (et EnhancedTransmission) ShiftDown() {
	fmt.Println("Enhanced Shift Down")
}
