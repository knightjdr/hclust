package cluster

import (
	"testing"

	"github.com/knightjdr/hclust/typedef"
	"github.com/stretchr/testify/assert"
)

func TestCluster(t *testing.T) {
	// TEST1: non-symmetric matrix.
	dist := [][]float64{
		{0, 10, 23, 22.6},
		{10, 0, 17.8, 17.4},
		{23, 17.8, 0, 12.2},
		{22.6, 17.4, 12.2, 0},
		{2, 5.8, 14.1, 15},
	}
	_, err := Cluster(dist, "single")
	assert.NotNil(t, err, "Non-symmetric matrix should return error")

	// TEST2: unknown linkage method.
	dist = [][]float64{
		{0, 10, 23, 22.6, 2},
		{10, 0, 17.8, 17.4, 5.8},
		{23, 17.8, 0, 12.2, 14.1},
		{22.6, 17.4, 12.2, 0, 15},
		{2, 5.8, 14.1, 15, 0},
	}
	_, err = Cluster(dist, "something")
	assert.NotNil(t, err, "Unknown linkage method should return error")

	// TEST3: single linkage.
	want := []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 5, Leafb: 1, Lengtha: 1.9, Lengthb: 2.9, Node: 6},
		{Leafa: 2, Leafb: 3, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 4.15, Lengthb: 0.95, Node: 8},
	}
	dendrogram, err := Cluster(dist, "single")
	assert.Nil(t, err, "Single linkage should not return an error")
	for i, cluster := range dendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Leaf a not added to dendrogram correctly for single linkage",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Leaf b not added to dendrogram correctly for single linkage",
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
		assert.Equal(
			t,
			want[i].Node,
			cluster.Node,
			"Parent node in subcluster not correct for single linkage",
		)
	}

	// Nearest Neighbor tests.
	dist = [][]float64{
		{0, 10, 23, 22.6, 2},
		{10, 0, 17.8, 17.4, 5.8},
		{23, 17.8, 0, 12.2, 14.1},
		{22.6, 17.4, 12.2, 0, 9},
		{2, 5.8, 14.1, 9, 0},
	}

	// TEST4: dendrogram for average method.
	want = []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 1, Leafb: 5, Lengtha: 3.95, Lengthb: 2.95, Node: 6},
		{Leafa: 2, Leafb: 3, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 4.71, Lengthb: 2.56, Node: 8},
	}
	dendrogram, _ = Cluster(dist, "average")
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
			"Parent node in subcluster not correct for average linkage",
		)
	}

	// TEST5: dendrogram for complete method.
	want = []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 1, Leafb: 5, Lengtha: 5, Lengthb: 4, Node: 6},
		{Leafa: 2, Leafb: 3, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 6.5, Lengthb: 5.4, Node: 8},
	}
	dendrogram, _ = Cluster(dist, "complete")
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
			"Parent node in subcluster not correct for complete linkage",
		)
	}

	// TEST6: dendrogram for mcquitty method.
	want = []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 1, Leafb: 5, Lengtha: 3.95, Lengthb: 2.95, Node: 6},
		{Leafa: 2, Leafb: 3, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 4.75, Lengthb: 2.59, Node: 8},
	}
	dendrogram, _ = Cluster(dist, "mcquitty")
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
			"Parent node in subcluster not correct for mcquitty linkage",
		)
	}

	// TEST7: dendrogram for ward method.
	want = []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 1, Leafb: 5, Lengtha: 4.68, Lengthb: 3.68, Node: 6},
		{Leafa: 2, Leafb: 3, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 8.06, Lengthb: 6.64, Node: 8},
	}
	dendrogram, _ = Cluster(dist, "ward")
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
			"Parent node in subcluster not correct for ward linkage",
		)
	}

	// Generic tests.
	dist = [][]float64{
		{0, 10, 23, 22.6, 2},
		{10, 0, 17.8, 17.4, 5.8},
		{23, 17.8, 0, 12.2, 14.1},
		{22.6, 17.4, 12.2, 0, 15},
		{2, 5.8, 14.1, 15, 0},
	}

	// TEST8: dendrogram for centroid method.
	want = []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 1, Leafb: 5, Lengtha: 4.06, Lengthb: 3.06, Node: 6},
		{Leafa: 2, Leafb: 3, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 4.53, Lengthb: 2.48, Node: 8},
	}
	dendrogram, _ = Cluster(dist, "centroid")
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
			"Parent node in subcluster not correct for centroid linkage",
		)
	}

	// TEST9: dendrogram for median method.
	want = []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 1, Leafb: 5, Lengtha: 4.06, Lengthb: 3.06, Node: 6},
		{Leafa: 2, Leafb: 3, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 4.36, Lengthb: 2.32, Node: 8},
	}
	dendrogram, _ = Cluster(dist, "median")
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
			"Parent node in subcluster not correct for median linkage",
		)
	}
}
