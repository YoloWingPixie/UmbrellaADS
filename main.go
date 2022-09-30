package main

import "fmt"

func main() {
	sam := matrix.create(0, 3)
	sam = matrix.addRow(matrix, []int{14, 2, 34})
	sam = matrix.addRow(matrix, []int{4, 54, 6})
	sam = matrix.addRow(matrix, []int{7, 8, 9})
	matrix.print(sam)

	fmt.Println()

	fmt.Println(findHighestInEachColumn(sam))
}
