package tree

import (
	"testing"

	"github.com/knightjdr/hclust/typedef"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	// TEST1: test a tree.
	dendrogram := []typedef.SubCluster{
		{Leafa: 0, Leafb: 3, Lengtha: 0.05, Lengthb: 0.05, Node: 6},
		{Leafa: 2, Leafb: 5, Lengtha: 0.075, Lengthb: 0.075, Node: 7},
		{Leafa: 1, Leafb: 6, Lengtha: 0.1, Lengthb: 0.05, Node: 8},
		{Leafa: 4, Leafb: 8, Lengtha: 0.2, Lengthb: 0.1, Node: 9},
		{Leafa: 7, Leafb: 9, Lengtha: 0.225, Lengthb: 0.1, Node: 10},
	}
	names := []string{"leaf0", "leaf1", "leaf2", "leaf3", "leaf4", "leaf5"}
	want := Level{
		Newick: "((leaf2:0.075,leaf5:0.075):0.225,(leaf4:0.2,(leaf1:0.1,(leaf0:0.05,leaf3:0.05):0.05):0.1):0.1)",
		Order:  []string{"leaf2", "leaf5", "leaf4", "leaf1", "leaf0", "leaf3"},
	}
	level := Create(dendrogram, names)
	assert.Equal(t, want.Order, level.Order, "Leafs not ordered correctly")
	assert.Equal(t, want.Newick, level.Newick, "Newick not formatted correctly")

	// TEST2: test another tree.
	dendrogram = []typedef.SubCluster{
		{Leafa: 0, Leafb: 4, Lengtha: 1, Lengthb: 1, Node: 5},
		{Leafa: 1, Leafb: 5, Lengtha: 3.95, Lengthb: 2.95, Node: 6},
		{Leafa: 2, Leafb: 3, Lengtha: 6.1, Lengthb: 6.1, Node: 7},
		{Leafa: 6, Leafb: 7, Lengtha: 4.71, Lengthb: 2.56, Node: 8},
	}
	names = []string{"leaf0", "leaf1", "leaf2", "leaf3", "leaf4"}
	want = Level{
		Newick: "((leaf1:3.95,(leaf0:1,leaf4:1):2.95):4.71,(leaf2:6.1,leaf3:6.1):2.56)",
		Order:  []string{"leaf1", "leaf0", "leaf4", "leaf2", "leaf3"},
	}
	level = Create(dendrogram, names)
	assert.Equal(t, want.Order, level.Order, "Leafs not ordered correctly")
	assert.Equal(t, want.Newick, level.Newick, "Newick not formatted correctly")
}
