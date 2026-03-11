package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(Relu(1.2))
	fmt.Println(Relu(-1))
	fmt.Println(Relu(7))
	fmt.Println(Relu(time.February))

	intMatrix, err := NewMatrix[int](2, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(intMatrix)

	intMatrix.Data[5] = 10
	fmt.Println(intMatrix.At(1, 2))             // 10
	fmt.Println(Max([]int{3, 1, 2}))            // 3
	fmt.Println(Max([]float64{3.5, 1, 2, 5.5})) // 5.5
	fmt.Println(Max([]float64(nil)))            // 0, err
}

type Matrix[T Number] struct {
	Rows int
	Cols int
	Data []T
}

func (m *Matrix[T]) At(rowIdx, colIdx int) T {
	i := (rowIdx * m.Cols) + colIdx

	return m.Data[i]
}

func Max[T Number](numbers []T) (T, error) {
	if len(numbers) == 0 {
		return *new(T), fmt.Errorf("no numbers provided")
	}

	max := numbers[0]
	for _, number := range numbers {
		if max < number {
			max = number
		}
	}

	return max, nil
}

func NewMatrix[T Number](rows, cols int) (*Matrix[T], error) {
	if rows < 0 || cols < 0 {
		return nil, fmt.Errorf("rows and cols can't be below zero, rows - %d, cols - %d", rows, cols)
	}

	return &Matrix[T]{
		Rows: rows,
		Cols: cols,
		Data: make([]T, rows*cols),
	}, nil
}

type Number interface {
	~int | ~float64
}

// T is type constraint, not type itself
func Relu[T Number](i T) T {
	if i < 0 {
		return 0
	}

	return i
}
