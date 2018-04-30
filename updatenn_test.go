package hclust

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateNN(t *testing.T) {
	dist := [][]float64{
		{0, 10, 23, 22.6, 2},
		{10, 0, 17.8, 17.4, 5.8},
		{23, 17.8, 0, 4.7, 3},
		{22.6, 17.4, 4.7, 0, 5.8},
		{2, 5.8, 3, 5.8, 0},
	}
	nodeSize := []int{2, 2, 1, 4, 1}

	// TEST1: unknown method.
	_, err := UpdateNN("unknown")
	assert.NotNil(t, err, "Unknown linkage method should return error")

	// TEST2: average.
	want := []float64{5, 5, 20.4, 20, 3.9, 0}
	updateFunc, _ := UpdateNN("average")
	assert.Equal(
		t,
		want,
		updateFunc(dist, 0, 1, nodeSize),
		"Dendrogram not correct for average linkage",
	)

	// TEST3: complete.
	want = []float64{10, 10, 23, 22.6, 5.8, 0}
	updateFunc, _ = UpdateNN("complete")
	assert.Equal(
		t,
		want,
		updateFunc(dist, 0, 1, nodeSize),
		"Dendrogram not correct for complete linkage",
	)

	// TEST4: mcquitty.
	want = []float64{5, 5, 20.4, 20, 3.9, 0}
	updateFunc, _ = UpdateNN("mcquitty")
	assert.Equal(
		t,
		want,
		updateFunc(dist, 0, 1, nodeSize),
		"Dendrogram not correct for mcquitty linkage",
	)

	// TEST5: ward.
	want = []float64{5.78, 5.78, 22.08, 23.67, 1.61, 0}
	updateFunc, _ = UpdateNN("ward")
	assert.InDeltaSlice(
		t,
		want,
		updateFunc(dist, 0, 1, nodeSize),
		0.01,
		"Dendrogram not correct for ward linkage",
	)
}
