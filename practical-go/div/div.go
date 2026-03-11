package main

import "fmt"

func main() {
	fmt.Println(div(7, 3))
	fmt.Println(safeDiv(7, 0))
}

func div(a, b int) int {
	return a / b
}

func safeDiv(a, b int) (q int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	return div(a, b), err
}
