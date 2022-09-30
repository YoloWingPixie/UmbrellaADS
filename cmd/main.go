package main

import "fmt"

func main() {

	matrix := createMatrix(0, 3)
	matrix = addRow(matrix, []int{14, 2, 34})
	matrix = addRow(matrix, []int{4, 54, 6})
	matrix = addRow(matrix, []int{7, 8, 9})
	printMatrix(matrix)

	fmt.Println()

	fmt.Println(findHighestInEachColumn(matrix))

}
