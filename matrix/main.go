package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rows := 90
	cols := 90
	minValue := 1
	maxValue := 12375091

	for {
		matrix := generateMatrix(rows, cols, minValue, maxValue)
		displayMatrix(matrix)

	}
}

func generateMatrix(rows, cols, minValue, maxValue int) [][]int {
	rand.Seed(time.Now().UnixNano())

	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = rand.Intn(maxValue-minValue+1) + minValue
		}
	}

	return matrix
}

func displayMatrix(matrix [][]int) {
	fmt.Println("Matrix:")
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%d\t", value)
		}
		fmt.Println()
	}
}
