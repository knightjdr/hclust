package hclust

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	matrix := [][]float64{
		{5, 2, 14.3, 2.1},
		{23, 17.8, 0, 0.4},
		{10, 0, 7, 15.9},
	}

	// TEST1: calculate distance between rows using maximum metric.
	want := [][]float64{
		{0, 18, 13.8},
		{18, 0, 17.8},
		{13.8, 17.8, 0},
	}
	dist := Distance(matrix, "maximum", false)
	// Iterate over rows to compare
	for i, row := range dist {
		assert.InDeltaSlice(t, want[i], row, 0.01, "Distance matrix not correct")
	}

	// TEST2: calculate distance between columns using maximum metric.
	want = [][]float64{
		{0, 10, 23, 22.6},
		{10, 0, 17.8, 17.4},
		{23, 17.8, 0, 12.2},
		{22.6, 17.4, 12.2, 0},
	}
	dist = Distance(matrix, "maximum", true)
	// Iterate over rows to compare
	for i, row := range dist {
		assert.InDeltaSlice(t, want[i], row, 0.01, "Distance of transposed matrix not correct")
	}
}
