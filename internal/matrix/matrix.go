package matrix

import (
	"fmt"
)

var ()

func Create(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}

func Print(matrix [][]int) {
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
}

// function to add a row to a matrix
func AddRow(matrix [][]int, row []int) [][]int {
	matrix = append(matrix, row)
	return matrix
}

func ColumnMaxMask(matrix [][]int) []int {
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

// function like ColumnMaxMask but it stops for each column once it finds a 1
func ColumnMaxMask2(matrix [][]int) []int {
	var result []int
	for i := 0; i < len(matrix[0]); i++ {
		var highest int
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] > highest {
				highest = matrix[j][i]
			}
			if highest == 1 {
				break
			}
		}
		result = append(result, highest)
	}
	return result
}
