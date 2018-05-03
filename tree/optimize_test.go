package tree

import (
	"testing"

	"github.com/knightjdr/hclust/typedef"
	"github.com/stretchr/testify/assert"
)

func TestOptimize(t *testing.T) {
	// TEST1: test a tree.
	dendrogram := []typedef.SubCluster{
		{Leafa: 0, Leafb: 3, Lengtha: 0.05, Lengthb: 0.05, Node: 6},
		{Leafa: 2, Leafb: 5, Lengtha: 0.075, Lengthb: 0.075, Node: 7},
		{Leafa: 1, Leafb: 6, Lengtha: 0.1, Lengthb: 0.05, Node: 8},
		{Leafa: 4, Leafb: 8, Lengtha: 0.2, Lengthb: 0.1, Node: 9},
		{Leafa: 7, Leafb: 9, Lengtha: 0.225, Lengthb: 0.1, Node: 10},
	}
	dist := [][]float64{
		{0, 0.21, 0.71, 0.1, 0.4, 0.72},
		{0.21, 0, 0.7, 0.2, 0.42, 0.71},
		{0.71, 0.7, 0, 0.72, 0.73, 0.125},
		{0.1, 0.2, 0.72, 0, 0.41, 0.73},
		{0.4, 0.42, 0.73, 0.41, 0, 0.74},
		{0.72, 0.71, 0.125, 0.73, 0.74, 0},
	}
	want := []typedef.SubCluster{
		{Leafa: 3, Leafb: 0, Lengtha: 0.05, Lengthb: 0.05, Node: 6},
		{Leafa: 5, Leafb: 2, Lengtha: 0.075, Lengthb: 0.075, Node: 7},
		{Leafa: 1, Leafb: 6, Lengtha: 0.1, Lengthb: 0.05, Node: 8},
		{Leafa: 8, Leafb: 4, Lengtha: 0.1, Lengthb: 0.2, Node: 9},
		{Leafa: 7, Leafb: 9, Lengtha: 0.225, Lengthb: 0.1, Node: 10},
	}
	optimized := Optimize(dendrogram, dist)
	assert.Equal(t, want, optimized, "Dendrogram not optimized correctly")
}
