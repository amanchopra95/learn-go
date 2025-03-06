package main

type startable interface {
	Start()
}

func startEngine(cars ...startable) {
	for _, car := range cars {
		car.Start()
	}
}

func main() {
	myConvertible := Convertible{Engine{}, EnhancedTransmission{}, SteeringWheel{}}
	myTruck := Truck{Engine{}, Transmission{}, SteeringWheel{}}

	myTruck.FourWheelDrive()
	myConvertible.ConvertTop()

	startEngine(myConvertible, myTruck)
}
