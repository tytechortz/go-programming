package main

import "fmt"

type burrito int

var x burrito

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}
