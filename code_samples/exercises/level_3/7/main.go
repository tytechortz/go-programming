package main

import "fmt"

func main() {
	x := 100
	y :=59
	if y > x {
		fmt.Println("greater")
	} else if y == x {
		fmt.Println("equal")
	} else {
		fmt.Println("less")
	}
}