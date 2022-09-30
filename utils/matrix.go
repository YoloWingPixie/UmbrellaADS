package utils

import (
	"fmt"
)

var ()

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
}

// function to add a row to a matrix
func addRow(matrix [][]int, row []int) [][]int {
	matrix = append(matrix, row)
	return matrix
}

func findHighestInEachColumn(matrix [][]int) []int {
	var result []int
	for i := 0; i < len(matrix[0]); i++ {
		var highest int
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] > highest {
				highest = matrix[j][i]
			}
		}
		result = append(result, highest)
	}
	return result
}
