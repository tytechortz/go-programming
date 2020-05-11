package main

import "fmt"

type person struct {
	first string
	last string
	favFlavors []string
}

func main() {
	p1 := person {
		first: "James",
		last: "Bond",
		favFlavors: []string{"chocolate", "vanilla",},
	}

	p2 := person {
		first: "Carol",
		last: "Comeau",
		favFlavors: []string{"mocha", "mint",},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}
	
}