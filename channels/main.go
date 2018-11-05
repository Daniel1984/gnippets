package main

import "fmt"

type carDetails struct {
	make       string
	nameLength int
}

func getCarInfoProducer(cars [5]string) <-chan carDetails {
	carDetailsChan := make(chan carDetails)

	go func() {
		defer close(carDetailsChan)
		for _, car := range cars {
			// channel must be closed from sender, not receiver
			carDetailsChan <- carDetails{car, len(car)}
		}
	}()

	return carDetailsChan
}

func main() {
	cars := [5]string{"bmw", "audi", "volkswagen", "volvo", "nissan"}
	carChan := getCarInfoProducer(cars)

	for carInfo := range carChan {
		fmt.Println(carInfo)
	}
}
