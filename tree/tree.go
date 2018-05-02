package tree

import "github.com/knightjdr/hclust/typedef"

// Tree references a tree in newick format and the leaf order.
type Tree struct {
	Newick string
	Order  []string
}

// Create generates a newick tree in string format and returns the order
// of the clustering.
func Create(dendrogram []typedef.SubCluster, names []string) (tree Tree) {
	// Dendrogram clusters/leaf number.
	n := len(dendrogram)

	// Create map of nodes to dendrogram indicies.
	nodeMap := make(map[int]int, n)
	for i, cluster := range dendrogram {
		nodeMap[cluster.Node] = i
	}

	// Begin with top node, iterate through left and right branches and add to
	// ordering.
	level := Descend(n, 2*n, nodeMap, dendrogram, names)
	tree.Newick = level.Newick
	tree.Order = level.Order
	return
}
