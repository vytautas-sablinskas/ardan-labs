package main

import "fmt"

func main() {
	var a any

	a = 7
	fmt.Println("a:", a)

	a = "Hi"
	fmt.Println("a", a)

	s, ok := a.(int)
	if !ok {
		fmt.Printf("failed to convert to int, actual type - %T", a)
	} else {
		fmt.Println("s", s)
	}
}
