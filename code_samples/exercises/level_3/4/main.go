package main

import "fmt"

func main() {
	x := 1972
	for {
		if x > 2020 {
			break
		}
		fmt.Println(x)
		x++
	}
}