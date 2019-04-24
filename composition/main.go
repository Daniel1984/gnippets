package main

import (
	"./models"
	"fmt"
)

type PersonPartOne struct {
	Surname string
	models.PersonPartTwo
}

func (p *PersonPartOne) SayHi() {
	fmt.Printf("Hello, my name is %s %s", p.Name, p.Surname)
}

func main() {
	p := &PersonPartOne{
		Surname: "Bar",
		PersonPartTwo: models.PersonPartTwo{
			Name: "Foo",
		},
	}

	p.SayHi()
}
