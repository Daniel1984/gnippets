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

type Aircraft struct {
	make     string
	topSpeed float32
}

func main() {
	bmw := &Car{make: "BMW", topSpeed: 312.5}
	audi := &Car{make: "Audi", topSpeed: 300}
	volga := &Car{make: "Gaz", topSpeed: 90}
	concorde := &Aircraft{make: "Concorde", topSpeed: 2405}
	drive(audi)
	drive(bmw)
	drive(volga)
	drive(concorde)
}

func (a Aircraft) getMake() string {
	return a.make
}

func (a Aircraft) drive(msg string) {
	fmt.Printf("Top speed of %v is %v kmh, %v \n", a.make, a.topSpeed, msg)
}

func (c Car) getMake() string {
	return c.make
}

func (c Car) drive(msg string) {
	fmt.Printf("Top speed of %v is %v kmh, %v \n", c.make, c.topSpeed, msg)
}

func logAction(msg string) {
	fmt.Printf("Performing logging for: %v \n", msg)
}

func drive(w Wehicle) {
	var msg string
	switch w.getMake() {
	case "BMW":
		msg = "preferred summer car"
	case "Audi":
		msg = "preferred winter car"
	case "Concorde":
		msg = "this is not a car"
	default:
		msg = "is this even a car?"
	}

	logAction(w.getMake())

	w.drive(msg)
}
