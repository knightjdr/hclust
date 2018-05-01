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
		{22.6, 17.4, 12.2, 0, 15},
		{2, 5.8, 14.1, 15, 0},
	}

	// TEST1: dendrogram.
	want := []SubCluster{
		{0, 4, 1, 1},
		{5, 1, 1.9, 2.9},
		{2, 3, 6.1, 6.1},
		{6, 7, 4.15, 0.95},
	}
	dendrogram := Single(dist)
	for i, cluster := range dendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Parent nodes not added to dendrogram correctly for single linkage",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Parent nodes not added to dendrogram correctly for single linkage",
		)
		assert.InDeltaf(
			t,
			want[i].Lengtha,
			cluster.Lengtha,
			0.01,
			"Dendrogram branch lengths not correct for single linkage",
		)
		assert.InDeltaf(
			t,
			want[i].Lengthb,
			cluster.Lengthb,
			0.01,
			"Dendrogram branch lengths not correct for single linkage",
		)
	}
}
