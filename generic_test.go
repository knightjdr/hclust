package hclust

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneric(t *testing.T) {
	dist := [][]float64{
		{0, 10, 23, 22.6, 2},
		{10, 0, 17.8, 17.4, 5.8},
		{23, 17.8, 0, 12.2, 14.1},
		{22.6, 17.4, 12.2, 0, 15},
		{2, 5.8, 14.1, 15, 0},
	}

	// TEST1: invalid method should return err.
	_, err := Generic(dist, "somemething")
	assert.NotNil(t, err, "Invalid method should return error")

	// TEST2: dendrogram for centroid method.
	want := []SubCluster{
		{0, 4, 1, 1},
		{1, 5, 4.06, 3.06},
		{2, 3, 6.1, 6.1},
		{6, 7, 4.53, 2.48},
	}
	dendrogram, _ := Generic(dist, "centroid")
	for i, cluster := range dendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Parent nodes not added to dendrogram correctly for centroid linkage",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Parent nodes not added to dendrogram correctly for centroid linkage",
		)
		assert.InDeltaf(
			t,
			want[i].Lengtha,
			cluster.Lengtha,
			0.01,
			"Dendrogram branch lengths not correct for centroid linkage",
		)
		assert.InDeltaf(
			t,
			want[i].Lengthb,
			cluster.Lengthb,
			0.01,
			"Dendrogram branch lengths not correct for centroid linkage",
		)
	}

	// TEST3: dendrogram for median method.
	want = []SubCluster{
		{0, 4, 1, 1},
		{1, 5, 4.06, 3.06},
		{2, 3, 6.1, 6.1},
		{6, 7, 4.36, 2.32},
	}
	dendrogram, _ = Generic(dist, "median")
	for i, cluster := range dendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Parent nodes not added to dendrogram correctly for median linkage",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Parent nodes not added to dendrogram correctly for median linkage",
		)
		assert.InDeltaf(
			t,
			want[i].Lengtha,
			cluster.Lengtha,
			0.01,
			"Dendrogram branch lengths not correct for median linkage",
		)
		assert.InDeltaf(
			t,
			want[i].Lengthb,
			cluster.Lengthb,
			0.01,
			"Dendrogram branch lengths not correct for median linkage",
		)
	}
}
