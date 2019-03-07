package main

import (
	"fmt"
)

type Order struct {
	Price    int
	Quantity int
}

func GetOrders() []*Order {
	orders := make([]*Order, 0)
	orders = append(orders, &Order{})
	orders = append(orders, &Order{})
	return orders
}

type FilterOne struct{}

func (d *FilterOne) Filter(input chan *Order) chan *Order {
	output := make(chan *Order)

	go func() {
		for o := range input {
			o.Price = 12
			output <- o
		}
		close(output)
	}()

	return output
}

type FilterTwo struct{}

func (d *FilterTwo) Filter(input chan *Order) chan *Order {
	output := make(chan *Order)
	go func() {
		for o := range input {
			o.Quantity = 2
			output <- o
		}
		close(output)
	}()
	return output
}

func main() {
	pipeline := InitPipeline(&FilterOne{}, &FilterTwo{})

	go func() {
		orders := GetOrders()
		for _, o := range orders {
			pipeline.Send(o)
		}
		pipeline.Close()
	}()

	pipeline.Receive(func(o *Order) {
		fmt.Printf("Received: %+v\n", o)
	})
}

type Pipeline struct {
	input  chan<- *Order
	output <-chan *Order
}

func (p *Pipeline) Send(o *Order) {
	p.input <- o
}

func (p *Pipeline) Receive(callback func(*Order)) {
	for o := range p.output {
		callback(o)
	}
}

func (p *Pipeline) Close() {
	close(p.input)
}

type Filter interface {
	Filter(input chan *Order) chan *Order
}

func InitPipeline(filters ...Filter) *Pipeline {
	source := make(chan *Order)
	var nextFilter chan *Order

	for _, filter := range filters {
		if nextFilter == nil {
			nextFilter = filter.Filter(source)
		} else {
			nextFilter = filter.Filter(nextFilter)
		}
	}

	return &Pipeline{input: source, output: nextFilter}
}
