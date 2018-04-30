package hclust

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNodes(t *testing.T) {
	dendrogram := []SubCluster{
		{0, 3, 0.1, 0.1},
		{2, 5, 0.15, 0.15},
		{1, 3, 0.2, 0.2},
		{4, 3, 0.4, 0.4},
		{5, 4, 0.6, 0.6},
	}

	// TEST1: add nodes to dendrogram.
	want := []SubCluster{
		{0, 3, 0.05, 0.05},
		{2, 5, 0.075, 0.075},
		{1, 6, 0.1, 0.05},
		{4, 8, 0.2, 0.1},
		{7, 9, 0.225, 0.1},
	}
	branchesDendrogram := AddNodes(dendrogram)
	for i, cluster := range branchesDendrogram {
		assert.Equal(
			t,
			want[i].Leafa,
			cluster.Leafa,
			"Parent nodes not added to dendrogram correctly",
		)
		assert.Equal(
			t,
			want[i].Leafb,
			cluster.Leafb,
			"Parent nodes not added to dendrogram correctly",
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
	}
}
