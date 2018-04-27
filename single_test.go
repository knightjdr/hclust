package hclust

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingle(t *testing.T) {
	dist := [][]float64{
		{0, 10, 23, 22.6, 2},
		{10, 0, 17.8, 17.4, 5.8},
		{23, 17.8, 0, 12.2, 14.1},
		{22.6, 17.4, 12.2, 0, 9},
		{2, 5.8, 14.1, 9, 0},
	}

	// TEST1: dendrogram.
	want := []SubCluster{
		{2, 0, 4},
		{5.8, 5, 1},
		{9, 6, 3},
		{12.2, 7, 2},
	}
	assert.Equal(t, want, Single(dist), "Dendrogram not correct for single linkage")
}
