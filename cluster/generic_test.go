package cluster

import (
	"testing"

	"github.com/knightjdr/hclust/typedef"
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
	want := []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 1, Leafb: 5, Lengtha: 4.06, Lengthb: 3.06, Node: 6},
		{Leafa: 2, Leafb: 3, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 4.53, Lengthb: 2.48, Node: 8},
	}
	dendrogram, _ := Generic(dist, "centroid")
	for i, cluster := range dendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Leaf a not added to dendrogram correctly for centroid linkage",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Leaf b not added to dendrogram correctly for centroid linkage",
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
		assert.Equal(
			t,
			want[i].Node,
			cluster.Node,
			"Parent node in subcluster not correct",
		)
	}

	// TEST3: dendrogram for median method.
	want = []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 1, Leafb: 5, Lengtha: 4.06, Lengthb: 3.06, Node: 6},
		{Leafa: 2, Leafb: 3, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 4.36, Lengthb: 2.32, Node: 8},
	}
	dendrogram, _ = Generic(dist, "median")
	for i, cluster := range dendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Leaf a not added to dendrogram correctly for median linkage",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Leaf b not added to dendrogram correctly for median linkage",
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
		assert.Equal(
			t,
			want[i].Node,
			cluster.Node,
			"Parent node in subcluster not correct",
		)
	}
}
