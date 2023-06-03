package main

import (
	"fmt"
	"math/rand"
	"time"
)

func multiplyMatrix(matrixA, matrixB [][]int) [][]int {
	rowsA := len(matrixA)
	colsA := len(matrixA[0])
	rowsB := len(matrixB)
	colsB := len(matrixB[0])

	if colsA != rowsB {
		panic("Invalid matrix dimensions. Cannot perform multiplication.")
	}

	result := make([][]int, rowsA)
	for i := 0; i < rowsA; i++ {
		result[i] = make([]int, colsB)
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				result[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}

	return result
}

func generateRandomMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = rand.Intn(10) // Generate random integers between 0 and 9
		}
	}

	return matrix
}

func main() {
	rowsA := 100
	colsA := 500
	rowsB := 500
	colsB := 200

	for {
		matrixA := generateRandomMatrix(rowsA, colsA)
		matrixB := generateRandomMatrix(rowsB, colsB)

		result := multiplyMatrix(matrixA, matrixB)

		fmt.Println("Matrix A:")
		printMatrix(matrixA)
		fmt.Println()

		fmt.Println("Matrix B:")
		printMatrix(matrixB)
		fmt.Println()

		fmt.Println("Matrix A * Matrix B:")
		printMatrix(result)
		fmt.Println()

		time.Sleep(1 * time.Second) // Delay between each matrix multiplication
	}
}

func printMatrix(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}
}
