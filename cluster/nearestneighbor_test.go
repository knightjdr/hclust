package cluster

import (
	"testing"

	"github.com/knightjdr/hclust/typedef"
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
	want := []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 1, Leafb: 5, Lengtha: 3.95, Lengthb: 2.95, Node: 6},
		{Leafa: 2, Leafb: 3, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 4.71, Lengthb: 2.56, Node: 8},
	}
	dendrogram, _ := NearestNeighbor(dist, "average")
	for i, cluster := range dendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Leaf a not added to dendrogram correctly for average linkage",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Leaf b not added to dendrogram correctly for average linkage",
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
		assert.Equal(
			t,
			want[i].Node,
			cluster.Node,
			"Parent node in subcluster not correct",
		)
	}

	// TEST3: dendrogram for complete method.
	want = []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 1, Leafb: 5, Lengtha: 5, Lengthb: 4, Node: 6},
		{Leafa: 2, Leafb: 3, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 6.5, Lengthb: 5.4, Node: 8},
	}
	dendrogram, _ = NearestNeighbor(dist, "complete")
	for i, cluster := range dendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Leaf a not added to dendrogram correctly for complete linkage",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Leaf b not added to dendrogram correctly for complete linkage",
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
		assert.Equal(
			t,
			want[i].Node,
			cluster.Node,
			"Parent node in subcluster not correct",
		)
	}

	// TEST4: dendrogram for mcquitty method.
	want = []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 1, Leafb: 5, Lengtha: 3.95, Lengthb: 2.95, Node: 6},
		{Leafa: 2, Leafb: 3, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 4.75, Lengthb: 2.59, Node: 8},
	}
	dendrogram, _ = NearestNeighbor(dist, "mcquitty")
	for i, cluster := range dendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Leaf a not added to dendrogram correctly for mcquitty linkage",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Leaf b not added to dendrogram correctly for mcquitty linkage",
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
		assert.Equal(
			t,
			want[i].Node,
			cluster.Node,
			"Parent node in subcluster not correct",
		)
	}

	// TEST5: dendrogram for ward method.
	want = []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 1, Leafb: 5, Lengtha: 4.68, Lengthb: 3.68, Node: 6},
		{Leafa: 2, Leafb: 3, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 8.06, Lengthb: 6.64, Node: 8},
	}
	dendrogram, _ = NearestNeighbor(dist, "ward")
	for i, cluster := range dendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Leaf a not added to dendrogram correctly for ward linkage",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Leaf b not added to dendrogram correctly for ward linkage",
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
		assert.Equal(
			t,
			want[i].Node,
			cluster.Node,
			"Parent node in subcluster not correct",
		)
	}
}
