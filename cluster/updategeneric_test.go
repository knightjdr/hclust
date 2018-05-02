package cluster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateGeneric(t *testing.T) {
	dist := [][]float64{
		{0, 10, 23, 22.6, 2},
		{10, 0, 17.8, 17.4, 5.8},
		{23, 17.8, 0, 4.7, 3},
		{22.6, 17.4, 4.7, 0, 5.8},
		{2, 5.8, 3, 5.8, 0},
	}
	nodeSize := []int{2, 2, 1, 4, 1}

	// TEST1: unknown method.
	_, err := UpdateGeneric("unknown")
	assert.NotNil(t, err, "Unknown linkage method should return error")

	// TEST2: centroid.
	want := []float64{17.55, 14.15, 3.43, 0.23, 3.71, 0}
	updateFunc, _ := UpdateGeneric("centroid")
	assert.InDeltaSlicef(
		t,
		want,
		updateFunc(dist, 3, 4, nodeSize),
		0.01,
		"Dendrogram not correct for centroid linkage",
	)

	// TEST3: median.
	want = []float64{10.85, 10.15, 2.40, 1.45, 1.45, 0}
	updateFunc, _ = UpdateGeneric("median")
	assert.InDeltaSlicef(
		t,
		want,
		updateFunc(dist, 3, 4, nodeSize),
		0.01,
		"Dendrogram not correct for median linkage",
	)
}
