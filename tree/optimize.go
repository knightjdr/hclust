package tree

import (
	"github.com/knightjdr/hclust/matrixop"
	"github.com/knightjdr/hclust/typedef"
)

type constraints struct {
	left  int
	right int
}

type leafs struct {
	a []int
	b []int
}

func optimal(n, node, leaf int, leafs []int, nodeScores map[int]float64, dist [][]float64) (score float64) {
	// Current best minimal score.
	score = 0
	if node == leaf {
		return
	} else {
		// Remove current leaf from leafs slice.
		leafIndex := matrixop.SliceIndex(len(leafs), func(i int) bool { return leafs[i] == leaf })
		availableLeafs := make([]int, len(leafs))
		copy(availableLeafs, leafs)
		availableLeafs = append(availableLeafs[:leafIndex], availableLeafs[leafIndex+1:]...)
		for _, leafb := range availableLeafs {
			if nodeScores[leafb] > score {
				score = nodeScores[leafb]
			}
		}
	}
	return
}

// Optimize optimizes the leafs ordering of a dendrogram using the method
// of Bar-Joseph, et al. 2001.
func Optimize(dendrogram []typedef.SubCluster, dist [][]float64) (optimized []typedef.SubCluster) {
	// Number of nodes and leafs.
	n := len(dendrogram)

	// Get leafs beneath each node and group them into two pools: leafs on the left (a)
	// go into one slice and nodes on the right (b) go into a second slice.
	nodeLeafs := make(map[int]leafs, n)
	for _, cluster := range dendrogram {
		// Get first group of leafs.
		aLeafs := make([]int, 0)
		if cluster.Leafa <= n { // If Leaf is a leaf.
			aLeafs = append(aLeafs, cluster.Leafa)
		} else { // If Leaf is a node.
			aLeafs = append(aLeafs, nodeLeafs[cluster.Leafa].a...)
			aLeafs = append(aLeafs, nodeLeafs[cluster.Leafa].b...)
		}
		// Get second group of leafs.
		bLeafs := make([]int, 0)
		if cluster.Leafb <= n {
			bLeafs = append(bLeafs, cluster.Leafb)
		} else {
			bLeafs = append(bLeafs, nodeLeafs[cluster.Leafb].a...)
			bLeafs = append(bLeafs, nodeLeafs[cluster.Leafb].b...)
		}
		nodeLeafs[cluster.Node] = leafs{a: aLeafs, b: bLeafs}
	}

	// Calculate optimal ordering score for each node.
	m := make(map[int]map[int]map[int]float64, n) // Optimal ordering map.
	for _, cluster := range dendrogram {
		node := cluster.Node
		leafs := append(nodeLeafs[node].a, nodeLeafs[node].b...)
		totalLeafs := len(leafs)
		// Initialize 2- and 3D of map.
		m[node] = make(map[int]map[int]float64, totalLeafs)
		for _, leaf := range leafs {
			m[node][leaf] = make(map[int]float64, totalLeafs)
		}
		// Iterate over leafs in pool a and compare against pool b.
		for _, aLeaf := range nodeLeafs[node].a {
			// Find optimal order for j as leftmost leaf.
			optScoreA := optimal(n, cluster.Leafa, aLeaf, nodeLeafs[node].a, m[cluster.Leafa][aLeaf], dist)
			for _, bLeaf := range nodeLeafs[node].b {
				// Find optimal order for k as rightmost leaf.
				optScoreB := optimal(n, cluster.Leafb, bLeaf, nodeLeafs[node].b, m[cluster.Leafb][bLeaf], dist)
				// Calculate score for current node with order j, k.
				optScore := optScoreA + optScoreB + dist[aLeaf][bLeaf]
				m[node][aLeaf][bLeaf] = optScore
				m[node][bLeaf][aLeaf] = optScore
			}
		}
	}

	// Re-order dendrogram.
	optimized = dendrogram

	// Constraints contains the left and right contraints for each node.
	constrain := make(map[int]constraints, n)
	constrain[optimized[n-1].Node] = constraints{left: -1, right: -1}

	// Iterate over nodes and reorder as needed
	for i := n - 1; i >= 0; i-- {
		node := optimized[i].Node
		// Find best leaf pair.
		maxDiff := float64(0)
		var outerA, outerB int
		if constrain[node].left >= 0 {
			for leafb, value := range m[node][constrain[node].left] {
				if value > maxDiff {
					maxDiff = value
					outerB = leafb
				}
			}
			outerA = constrain[node].left
		} else if constrain[node].right >= 0 {
			for leafa, value := range m[node][constrain[node].right] {
				if value > maxDiff {
					maxDiff = value
					outerA = leafa
				}
			}
			outerB = constrain[node].right
		} else {
			for leafa := range m[node] {
				for leafb, value := range m[node][leafa] {
					if value > maxDiff {
						maxDiff = value
						outerA = leafa
						outerB = leafb
					}
				}
			}
		}

		// Check if outerA leaf is already in left pool, if not switch left and
		// right leafs.
		leafAIndex := matrixop.SliceIndex(len(nodeLeafs[node].a), func(j int) bool { return nodeLeafs[node].a[j] == outerA })
		if leafAIndex < 0 {
			leafa := optimized[i].Leafa
			leafb := optimized[i].Leafb
			lengtha := optimized[i].Lengtha
			lengthb := optimized[i].Lengthb
			optimized[i].Leafa = leafb
			optimized[i].Leafb = leafa
			optimized[i].Lengtha = lengthb
			optimized[i].Lengthb = lengtha
		}

		// Set contraints for subnodes.
		if optimized[i].Leafa > n {
			constrain[optimized[i].Leafa] = constraints{left: outerA, right: -1}
		}
		if optimized[i].Leafb > n {
			constrain[optimized[i].Leafb] = constraints{left: -1, right: outerB}
		}
	}

	return
}
