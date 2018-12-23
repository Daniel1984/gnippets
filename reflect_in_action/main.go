package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X     int
	Y     int
	Label string
}

func main() {
	p1 := &Point{20, 30, "point 1"}
	p2 := &Point{Label: "point 2"}

	s1 := reflect.ValueOf(p1).Elem()
	s2 := reflect.ValueOf(p2).Elem()

	fmt.Println("s2 =", s2)

	typeOfS1 := s1.Type()
	fmt.Println("p1 =", p1)
	fmt.Println("p2 =", p2)

	for i := 0; i < s1.NumField(); i++ {
		f := s1.Field(i)
		fmt.Printf("%d: %s ", i, typeOfS1.Field(i).Name)
		fmt.Printf("%s = %v\n", f.Type(), f.Interface())
	}
}
