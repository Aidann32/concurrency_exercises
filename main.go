package main

import (
	"fmt"
	"github.com/Aidann32/concurrency_exercises/exercises/files"
	"github.com/Aidann32/concurrency_exercises/exercises/matrix"
)

func main() {
	files.Run(10)

	matrixA := [][]int{
		{1, 2},
		{3, 4},
	}
	matrixB := [][]int{
		{5, 6},
		{7, 8},
	}
	multiplication := matrix.NewMultiplication(matrixA, matrixB, 10)
	fmt.Println(multiplication.Run())

	matrixA = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	matrixB = [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}
	multiplication = matrix.NewMultiplication(matrixA, matrixB, 10)
	fmt.Println(multiplication.Run())

	matrixA = [][]int{
		{2, 0, -1},
		{3, 5, 2},
		{1, 4, 3},
	}
	matrixB = [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}
	multiplication = matrix.NewMultiplication(matrixA, matrixB, 10)
	fmt.Println(multiplication.Run())
}
