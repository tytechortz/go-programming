package main

import "fmt"

func main() {
	x := [5]int{1, 2, 4, 8, 16}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}