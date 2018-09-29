package main

import "fmt"

type carDetails struct {
	make       string
	nameLength int
}

func calculateCarnamesLengths(cars [5]string, c chan carDetails) {
	for _, car := range cars {
		nameLength := len(car)
		c <- carDetails{car, nameLength}
	}

	// channel must be closed from sender, not receiver
	close(c)
}

func main() {
	cars := [5]string{"bmw", "audi", "volkswagen", "volvo", "nissan"}
	carChan := make(chan carDetails)
	go calculateCarnamesLengths(cars, carChan)

	for carInfo := range carChan {
		fmt.Println(carInfo)
	}
}
