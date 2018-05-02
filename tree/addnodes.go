package tree

import "github.com/knightjdr/hclust/typedef"

type union struct {
	Length    []float64
	NextLabel int
	Parent    []int
}

// Find highest incorporated parent node for a subnode n.
func (u *union) Find(n int, length float64) (int, float64) {
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

	// Get length of parent or set length to half input value if not defined.
	nodeLength := length / float64(2)
	if u.Length[n] > -1 {
		nodeLength = nodeLength - u.Length[n]
	}
	return n, nodeLength
}

// Set parent of most recently added node. Also set it's length.
func (u *union) AddParent(a, b int, length float64) {
	u.Length[u.NextLabel] = length
	u.Parent[a] = u.NextLabel
	u.Parent[b] = u.NextLabel
	u.NextLabel++
}

// AddNodes adds numbered nodes to a dendrogram and converts distances between
// leafs to branch lengths. The first new node will be equal to the length of
// the dendrogram.
func AddNodes(dendrogram []typedef.SubCluster) (labelledDendrogram []typedef.SubCluster) {
	// First parent node number.
	n := len(dendrogram) + 1

	// Create labels. Unknown parent nodes and node lengths are -1, i.e. not known.
	length := make([]float64, 2*len(dendrogram)+1)
	parent := make([]int, 2*len(dendrogram))
	for i := range parent {
		length[i] = -1
		parent[i] = -1
	}
	labels := union{Length: length, NextLabel: n, Parent: parent}

	// First node to add.
	for _, subcluster := range dendrogram {
		subnodeA, lengthA := labels.Find(subcluster.Leafa, subcluster.Lengtha)
		subnodeB, lengthB := labels.Find(subcluster.Leafb, subcluster.Lengthb)
		labelledDendrogram = append(
			labelledDendrogram,
			typedef.SubCluster{
				Leafa:   subnodeA,
				Leafb:   subnodeB,
				Lengtha: lengthA,
				Lengthb: lengthB,
				Node:    labels.NextLabel,
			},
		)
		labels.AddParent(subnodeA, subnodeB, subcluster.Lengtha/float64(2))
	}
	return
}
