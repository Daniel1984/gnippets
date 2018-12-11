package main

import "fmt"

func main() {
	carsSlice := []string{"Audi", "Volvo", "Bmw", "Bmw"}
	uniqueCarsSlice := unique(carsSlice)
	fmt.Println(uniqueCarsSlice)
}

func unique(slice []string) []string {
	foundValues := make(map[string]bool)
	uniqueElements := []string{}

	for _, entry := range slice {
		if _, value := foundValues[entry]; !value {
			foundValues[entry] = true
			uniqueElements = append(uniqueElements, entry)
		}
	}

	return uniqueElements
}
