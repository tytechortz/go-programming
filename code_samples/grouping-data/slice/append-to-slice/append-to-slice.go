package main

import "fmt"

func main() {
	x := []int{4,5,7,8,42}
	x = append(x, 77, 88, 99, 1014)
	y :=[]int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)
}