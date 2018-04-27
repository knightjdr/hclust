package hclust

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNodes(t *testing.T) {
	dendrogram := []SubCluster{
		{0.1, 0, 3},
		{0.15, 2, 5},
		{0.2, 1, 3},
		{0.4, 4, 3},
		{0.6, 5, 4},
	}

	// TEST1: add nodes to dendrogram.
	want := []SubCluster{
		{0.1, 0, 3},
		{0.15, 2, 5},
		{0.2, 1, 6},
		{0.4, 4, 8},
		{0.6, 7, 9},
	}
	assert.Equal(t, want, AddNodes(dendrogram), "Parent nodes not added to dendrogram correctly")
}
