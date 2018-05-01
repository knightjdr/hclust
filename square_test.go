package hclust

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare(t *testing.T) {
	dist := [][]float64{
		{0, 10, 23, 20, 2},
		{10, 0, 17.8, 17.4, 5.8},
		{23, 17.8, 0, 12.2, 14.1},
		{20, 17.4, 12.2, 0, 15},
		{2, 5.8, 14.1, 15, 0},
	}

	// TEST1: square a matrix.
	want := [][]float64{
		{0, 100, 529, 400, 4},
		{100, 0, 316.84, 302.76, 33.64},
		{529, 316.84, 0, 148.84, 198.81},
		{400, 302.76, 148.84, 0, 225},
		{4, 33.64, 198.81, 225, 0},
	}
	matrix := Square(dist)
	for i, row := range matrix {
		assert.InDeltaSlicef(
			t,
			want[i],
			row,
			0.01,
			"Matrix not squared correctly",
		)
	}
}
