package tree

import (
	"testing"

	"github.com/knightjdr/hclust/typedef"
	"github.com/stretchr/testify/assert"
)

func TestAddNodes(t *testing.T) {
	dendrogram := []typedef.SubCluster{
		{Leafa: 0, Leafb: 3, Lengtha: 0.1, Lengthb: 0.1, Node: 0},
		{Leafa: 2, Leafb: 5, Lengtha: 0.15, Lengthb: 0.15, Node: 0},
		{Leafa: 1, Leafb: 3, Lengtha: 0.2, Lengthb: 0.2, Node: 0},
		{Leafa: 4, Leafb: 3, Lengtha: 0.4, Lengthb: 0.4, Node: 0},
		{Leafa: 5, Leafb: 4, Lengtha: 0.6, Lengthb: 0.6, Node: 0},
	}

	// TEST1: add nodes to dendrogram.
	want := []typedef.SubCluster{
		{Leafa: 0, Leafb: 3, Lengtha: 0.05, Lengthb: 0.05, Node: 6},
		{Leafa: 2, Leafb: 5, Lengtha: 0.075, Lengthb: 0.075, Node: 7},
		{Leafa: 1, Leafb: 6, Lengtha: 0.1, Lengthb: 0.05, Node: 8},
		{Leafa: 4, Leafb: 8, Lengtha: 0.2, Lengthb: 0.1, Node: 9},
		{Leafa: 7, Leafb: 9, Lengtha: 0.225, Lengthb: 0.1, Node: 10},
	}
	branchesDendrogram := AddNodes(dendrogram)
	for i, cluster := range branchesDendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Leaf a not added to dendrogram correctly",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Leaf b not added to dendrogram correctly",
		)
		assert.InDeltaf(
			t,
			want[i].Lengtha,
			cluster.Lengtha,
			0.01,
			"Dendrogram branch lengths not correct",
		)
		assert.InDeltaf(
			t,
			want[i].Lengthb,
			cluster.Lengthb,
			0.01,
			"Dendrogram branch lengths not correct",
		)
		assert.Equal(
			t,
			want[i].Node,
			cluster.Node,
			"Parent node in subcluster not correct",
		)
	}
}
