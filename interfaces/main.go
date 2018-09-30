package main

import "fmt"

type Wehicle interface {
	getMake() string
	drive(string)
}

type Car struct {
	make     string
	topSpeed float32
}

func main() {
	bmw := &Car{make: "BMW", topSpeed: 312.5}
	audi := &Car{make: "Audi", topSpeed: 300}
	volga := &Car{make: "Gaz", topSpeed: 90}
	driveCar(audi)
	driveCar(bmw)
	driveCar(volga)
}

func (c Car) getMake() string {
	return c.make
}

func (c Car) drive(msg string) {
	fmt.Printf("Top speed of %v is %v kmh, %v \n", c.make, c.topSpeed, msg)
}

func driveCar(w Wehicle) {
	var msg string
	switch w.getMake() {
	case "BMW":
		msg = "preferred summer car"
	case "Audi":
		msg = "preferred winter car"
	default:
		msg = "is this even a car?"
	}

	w.drive(msg)
}
