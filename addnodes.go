package hclust

type Union struct {
	NextLabel int
	Parent    []int
}

// Find highest incorporated parent node for a subnode n.
func (u *Union) Find(n int) int {
	p := n

	// Find existing highest parent in dendrogram.
	for u.Parent[n] >= 0 {
		n = u.Parent[n]
	}

	// Redefine parent of node n (and it's parents) to highest available.
	// This makes subsequent searches for top parent faster.
	for p > -1 && u.Parent[p] != n {
		currP := p
		p = u.Parent[currP]
		u.Parent[currP] = n
	}
	return n
}

// Set parent of most recently added node.
func (u *Union) AddParent(a, b int) {
	u.Parent[a] = u.NextLabel
	u.Parent[b] = u.NextLabel
	u.NextLabel++
}

// AddNodes adds numbered nodes to a dendrogram. The first new node will be
// equal to the length of the dendrogram.
func AddNodes(dendrogram []SubCluster) (newDendrogram []SubCluster) {
	// First parent node number.
	N := len(dendrogram) + 1

	// Create union. Unknown parent nodes are -1.
	parent := make([]int, 2*len(dendrogram))
	for i := range parent {
		parent[i] = -1
	}
	union := Union{NextLabel: N, Parent: parent}

	// First node to add.
	for _, subcluster := range dendrogram {
		subnodeA := union.Find(subcluster.Leafa)
		subnodeB := union.Find(subcluster.Leafb)
		newDendrogram = append(newDendrogram, SubCluster{subcluster.Dist, subnodeA, subnodeB})
		union.AddParent(subnodeA, subnodeB)
	}
	return
}
