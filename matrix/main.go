package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rows := 500
	cols := 500

	symbols := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+{}[]|\\;:'\",.<>?")

	for {
		matrix := generateMatrix(rows, cols, symbols)
		displayMatrix(matrix)

	}
}

func generateMatrix(rows, cols int, symbols []rune) [][]rune {
	rand.Seed(time.Now().UnixNano())

	matrix := make([][]rune, rows)
	flattened := make([]rune, rows*cols)

	for i := 0; i < rows; i++ {
		matrix[i] = flattened[i*cols : (i+1)*cols]
		for j := 0; j < cols; j++ {
			matrix[i][j] = symbols[rand.Intn(len(symbols))]
		}
	}

	return matrix
}

func displayMatrix(matrix [][]rune) {
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("\033[32m%c\t\033[0m", value) // Set text color to green
		}
		fmt.Println()
	}
}
