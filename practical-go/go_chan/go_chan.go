package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("test")
	fmt.Println("main")

	for i := range 3 {
		go func() {
			fmt.Println("goroutine", i)
		}()
	}

	ch := make(chan int)
	go func() {
		ch <- 7 // send
	}()

	v := <-ch
	fmt.Println(v)

	fmt.Println(sleepSort([]int64{20, 30, 10})) // [10, 20, 30]
}

func sleepSort(numbers []int64) []int64 {
	sorted := make([]int64, 0, len(numbers))
	ch := make(chan int64)

	for _, number := range numbers {
		go func() {
			time.Sleep(time.Duration(number) * time.Millisecond)
			ch <- number
		}()
	}

	for _ = range len(numbers) {
		sorted = append(sorted, <-ch)
	}

	return sorted
}
