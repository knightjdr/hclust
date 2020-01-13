package dendrogram

import (
	"testing"

	"github.com/knightjdr/hclust/typedef"
	"github.com/stretchr/testify/assert"
)

func TestGetNodeHeight(t *testing.T) {
	// TEST1: test a tree.
	dendrogram := []typedef.SubCluster{
		{Leafa: 0, Leafb: 3, Lengtha: 0.05, Lengthb: 0.05, Node: 6},
		{Leafa: 2, Leafb: 5, Lengtha: 0.075, Lengthb: 0.075, Node: 7},
		{Leafa: 1, Leafb: 6, Lengtha: 0.1, Lengthb: 0.05, Node: 8},
		{Leafa: 4, Leafb: 8, Lengtha: 0.2, Lengthb: 0.1, Node: 9},
		{Leafa: 7, Leafb: 9, Lengtha: 0.225, Lengthb: 0.1, Node: 10},
	}
	want := []float64{0.1, 0.15, 0.25, 0.55, 1.025}
	assert.Equal(t, want, GetNodeHeight(dendrogram), "Should calculate node heights")

}
