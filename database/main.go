package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"log"
)

type car struct {
	Id    int
	Make  string
	Model string
}

func main() {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Database: "gnippets",
	})

	defer db.Close()

	// log.Fatal terminates process so else conditios below are redundant
	// but I personally think this is slighly more readable

	var n int
	if _, err := db.QueryOne(pg.Scan(&n), "SELECT 1"); err != nil {
		log.Fatal("Error establishing connection to db:", err)
	} else {
		fmt.Println("db connected!:", db)
	}

	// create new car
	if err := db.Insert(&car{Make: "Ford", Model: "Mustang"}); err != nil {
		log.Fatal("Error creating car:", err)
	} else {
		fmt.Println("Record created")
	}

	// select one car
	someCar := &car{Id: 1}
	if err := db.Select(someCar); err != nil {
		log.Fatal("Where is my cars dude?", err)
	} else {
		fmt.Printf("Got car, sweet: %v \n", someCar)
	}

	// select all cars
	var cars []car
	if err := db.Model(&cars).Select(); err != nil {
		log.Fatal("Where are my cars dude?", err)
	} else {
		fmt.Printf("Got cars, sweet: %v\n", cars)
	}

	// select specific cars
	var specificCars []car
	if err := db.Model(&specificCars).Where("make = ?", "Ford").Select(); err != nil {
		log.Fatal("Can't query cars:", err)
	} else {
		fmt.Printf("Specific cars: %v\n", specificCars)
	}

	// updating car
	carToUpdate := &car{Id: 1}
	if err := db.Select(carToUpdate); err != nil {
		log.Fatal("Can't find car", err)
	}

	carToUpdate.Model = "R8"
	if err := db.Update(carToUpdate); err != nil {
		log.Fatal("Couldn't update car", err)
	}

	if err := db.Select(carToUpdate); err != nil {
		log.Fatal("Couldn't find a car", err)
	} else {
		fmt.Printf("Car updated to: %v", carToUpdate)
	}
}
