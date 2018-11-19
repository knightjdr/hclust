package optimize

import (
	"testing"

	"github.com/knightjdr/hclust/typedef"
	"github.com/stretchr/testify/assert"
)

func TestOptimize(t *testing.T) {
	// TEST: find maximum between two integers
	max := maxInt(10, 8)
	assert.Equal(t, 10, max, "Maximum integer not returned")

	// TEST: find minimum between two integers
	min := minInt(10, 8)
	assert.Equal(t, 8, min, "Minimum integer not returned")

	// TEST: sort a map based on keys
	testMap := map[int]float64{
		1: 10.4,
		2: 4.3,
		4: 75.4,
	}
	expectedOrder := []int{2, 1, 4}
	acutalOrder := sortMap(testMap)
	assert.Equal(t, expectedOrder, acutalOrder, "Map not sorted correctly by keys")

	// TEST: shouldIgnore should always return false with ignore is "0"
	ignoreFunc := shouldIgnore(0)
	assert.False(t, ignoreFunc(0), "Should not ignore node with 0 comparisons when ignore arg is 0")
	assert.False(t, ignoreFunc(1000000), "Should not ignore node with 1M comparisons when ignore arg is 0")

	// TEST: shouldIgnore when ignore is > 0
	ignoreFunc = shouldIgnore(250000)
	assert.False(t, ignoreFunc(100000), "Should not ignore node with 100K comparisons when ignore arg is 250K")
	assert.True(t, ignoreFunc(250000), "Should ignore node with 250K comparisons when ignore arg is 250K")

	// TEST: test a tree.
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
		{Leafa: 0, Leafb: 3, Lengtha: 0.05, Lengthb: 0.05, Node: 6},
		{Leafa: 2, Leafb: 5, Lengtha: 0.075, Lengthb: 0.075, Node: 7},
		{Leafa: 6, Leafb: 1, Lengtha: 0.05, Lengthb: 0.1, Node: 8},
		{Leafa: 4, Leafb: 8, Lengtha: 0.2, Lengthb: 0.1, Node: 9},
		{Leafa: 9, Leafb: 7, Lengtha: 0.1, Lengthb: 0.225, Node: 10},
	}
	optimized := Optimize(dendrogram, dist, 0)
	assert.Equal(t, want, optimized, "Dendrogram not optimized correctly")

	// TEST: test another tree.
	dendrogram = []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 1, Leafb: 5, Lengtha: 3.95, Lengthb: 2.95, Node: 6},
		{Leafa: 2, Leafb: 3, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 4.71, Lengthb: 2.56, Node: 8},
	}
	dist = [][]float64{
		{0, 7.91, 17.7, 17.6, 2},
		{7.91, 0, 17.33, 17.32, 7.9},
		{17.7, 17.33, 0, 12.2, 17.5},
		{17.6, 17.32, 12.2, 0, 17.4},
		{2, 7.9, 17.7, 17.4, 0},
	}
	want = []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 5, Leafb: 1, Lengtha: 2.95, Lengthb: 3.95, Node: 6},
		{Leafa: 3, Leafb: 2, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 4.71, Lengthb: 2.56, Node: 8},
	}
	optimized = Optimize(dendrogram, dist, 0)
	assert.Equal(t, want, optimized, "Dendrogram not optimized correctly")
}
