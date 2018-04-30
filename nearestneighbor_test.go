package hclust

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNearestNeighbor(t *testing.T) {
	dist := [][]float64{
		{0, 10, 23, 22.6, 2},
		{10, 0, 17.8, 17.4, 5.8},
		{23, 17.8, 0, 12.2, 14.1},
		{22.6, 17.4, 12.2, 0, 9},
		{2, 5.8, 14.1, 9, 0},
	}

	// TEST1: invalid method should return err.
	_, err := NearestNeighbor(dist, "somemething")
	assert.NotNil(t, err, "Invalid method should return error")

	// TEST2: dendrogram for average method.
	want := []SubCluster{
		{0, 4, 1, 1},
		{1, 5, 3.95, 2.95},
		{2, 3, 6.1, 6.1},
		{6, 7, 4.71, 2.56},
	}
	dendrogram, _ := NearestNeighbor(dist, "average")
	for i, cluster := range dendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Parent nodes not added to dendrogram correctly for average linkage",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Parent nodes not added to dendrogram correctly for average linkage",
		)
		assert.InDeltaf(
			t,
			want[i].Lengtha,
			cluster.Lengtha,
			0.01,
			"Dendrogram branch lengths not correct for average linkage",
		)
		assert.InDeltaf(
			t,
			want[i].Lengthb,
			cluster.Lengthb,
			0.01,
			"Dendrogram branch lengths not correct for average linkage",
		)
	}

	// TEST3: dendrogram for complete method.
	want = []SubCluster{
		{0, 4, 1, 1},
		{1, 5, 5, 4},
		{2, 3, 6.1, 6.1},
		{6, 7, 6.5, 5.4},
	}
	dendrogram, _ = NearestNeighbor(dist, "complete")
	for i, cluster := range dendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Parent nodes not added to dendrogram correctly for complete linkage",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Parent nodes not added to dendrogram correctly for complete linkage",
		)
		assert.InDeltaf(
			t,
			want[i].Lengtha,
			cluster.Lengtha,
			0.01,
			"Dendrogram branch lengths not correct for complete linkage",
		)
		assert.InDeltaf(
			t,
			want[i].Lengthb,
			cluster.Lengthb,
			0.01,
			"Dendrogram branch lengths not correct for complete linkage",
		)
	}

	// TEST4: dendrogram for mcquitty method.
	want = []SubCluster{
		{0, 4, 1, 1},
		{1, 5, 3.95, 2.95},
		{2, 3, 6.1, 6.1},
		{6, 7, 4.75, 2.59},
	}
	dendrogram, _ = NearestNeighbor(dist, "mcquitty")
	for i, cluster := range dendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Parent nodes not added to dendrogram correctly for mcquitty linkage",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Parent nodes not added to dendrogram correctly for mcquitty linkage",
		)
		assert.InDeltaf(
			t,
			want[i].Lengtha,
			cluster.Lengtha,
			0.01,
			"Dendrogram branch lengths not correct for mcquitty linkage",
		)
		assert.InDeltaf(
			t,
			want[i].Lengthb,
			cluster.Lengthb,
			0.01,
			"Dendrogram branch lengths not correct for mcquitty linkage",
		)
	}

	// TEST5: dendrogram for ward method.
	want = []SubCluster{
		{0, 4, 1, 1},
		{1, 5, 4.68, 3.68},
		{2, 3, 6.1, 6.1},
		{6, 7, 8.06, 6.64},
	}
	dendrogram, _ = NearestNeighbor(dist, "ward")
	for i, cluster := range dendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Parent nodes not added to dendrogram correctly for ward linkage",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Parent nodes not added to dendrogram correctly for ward linkage",
		)
		assert.InDeltaf(
			t,
			want[i].Lengtha,
			cluster.Lengtha,
			0.01,
			"Dendrogram branch lengths not correct for ward linkage",
		)
		assert.InDeltaf(
			t,
			want[i].Lengthb,
			cluster.Lengthb,
			0.01,
			"Dendrogram branch lengths not correct for ward linkage",
		)
	}
}
