package main

import "fmt"

func main() {
	ii := []int{2, 3, 4, 5, 6, 7,}
	n := foo(ii...)
	fmt.Println(n)

	ii2 := []int{2, 3, 4, 5, 6, 7,}
	n2 := bar(ii2)
	fmt.Println(n2)
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}