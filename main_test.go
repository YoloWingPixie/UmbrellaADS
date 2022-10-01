package main

import (
	"math/rand"
	"testing"
	matrix "umbrella/internal/matrix"
)

// benchmarking the two functions
func BenchmarkColumnMaxMask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sam := matrix.Create(0, 10)

		//populate the matrix with 100 rows of 10 columns containing either 0 or 1
		for k := 0; k < 100; k++ {
			var row []int
			for l := 0; l < 10; l++ {
				row = append(row, rand.Intn(2))
			}
			sam = matrix.AddRow(sam, row)
		}

		// randomly select 30 rows from the matrix into a new matrix
		var newMatrix [][]int
		for i := 0; i < 30; i++ {
			newMatrix = matrix.AddRow(newMatrix, sam[rand.Intn(len(sam))])
		}

		matrix.ColumnMaxMask2(newMatrix)
	}
}
