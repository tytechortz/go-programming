package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age)
}

func main() {
	p1 := person{
		first: "Chris",
		last:"Hedges",
		age: 45,
	}

	p1.speak()
}