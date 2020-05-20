package main 

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("This should run last")
}

func bar() {
	fmt.Println("This should run first")
}