package main

import (
	"fmt"
	"slices"
)

func main() {
	cart := []string{"apple", "orange", "banana"}
	fmt.Println(len(cart))
	fmt.Println(cart[1])

	for i, item := range cart {
		fmt.Println(i, item)
	}

	cart = append(cart, "milk")
	fmt.Println(cart)

	fruits := cart[:3]
	fmt.Println(fruits)
	fruits = append(fruits, "lemon")
	fmt.Println("fruits:", fruits)
	fmt.Println("cart:", cart)

	// Exercise
	fmt.Println(concat([]string{"A", "B"}, []string{"C"})) // expected: [A B C]

	fmt.Println(median([]float64{3, 1, 2}))    // expected 2
	fmt.Println(median([]float64{3, 1, 2, 4})) // expected 2.5
}

func concat(s1, s2 []string) []string {
	s3 := make([]string, len(s1)+len(s2))
	copy(s3, s1)
	copy(s3[len(s1):], s2)

	return s3
}

func median(values []float64) float64 {
	slices.Sort(values)

	if len(values)%2 == 0 {
		return (values[len(values)/2-1] + values[len(values)/2]) / 2
	}

	return values[len(values)/2]
}
