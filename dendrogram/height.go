// Package dendrogram has helpers for analyzing dendrograms.
package dendrogram

import (
	"github.com/knightjdr/hclust/typedef"
)

// GetNodeHeight gets the height for each node by summing branch lengths.
func GetNodeHeight(dendrogram []typedef.SubCluster) []float64 {
	nodeHeight := make([]float64, len(dendrogram))
	startingNode := len(dendrogram) + 1

	for nodeIndex, node := range dendrogram {
		nodeHeight[nodeIndex] = sumBranchLengths(node, nodeHeight, startingNode)
	}

	return nodeHeight
}

func sumBranchLengths(node typedef.SubCluster, nodeHeight []float64, startingNode int) float64 {
	height := node.Lengtha + node.Lengthb
	if node.Leafa >= startingNode {
		height += nodeHeight[node.Leafa-startingNode]
	}
	if node.Leafb >= startingNode {
		height += nodeHeight[node.Leafb-startingNode]
	}

	return height
}
