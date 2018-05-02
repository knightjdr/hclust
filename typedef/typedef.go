// Package typedef has type definitions used throught the clust package.
package typedef

// Hclust contains the dendrogram, newick tree and cluster order.
type Hclust struct {
	Dendrogram []SubCluster
	Newick     string
	Order      []string
}

// SubCluster stores the node, distance and names of leafs for a subcluster.
type SubCluster struct {
	Leafa   int
	Leafb   int
	Lengtha float64
	Lengthb float64
	Node    int
}
