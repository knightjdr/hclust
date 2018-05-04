// Package optimize optimizes the clustering order.
package optimize

import (
	"math"
	"sort"

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

// Optimal implements the "fast" leaf optimization approach of Bar-Joseph et al.
// 2001. See Figure 4.
func optimal(aSortOrder, bSortOrder []int, minDist float64, nodeScoresA map[int]float64, nodeScoresB map[int]float64, dist [][]float64) (score float64) {
	// Current best maximal score.
	score = math.MaxFloat64
	for _, leftIndex := range aSortOrder {
		ma := nodeScoresA[leftIndex]
		if ma+nodeScoresB[bSortOrder[0]]+minDist >= score {
			return
		}
		for _, rightIndex := range bSortOrder {
			mb := nodeScoresB[rightIndex]
			if ma+mb+minDist >= score {
				break
			} else if score > ma+mb+dist[leftIndex][rightIndex] {
				score = ma + mb + dist[leftIndex][rightIndex]
			}
		}
	}
	return
}

// SortMap sorts a map in descending order based on its keys.
func sortMap(unsortedMap map[int]float64) (sortOrder []int) {
	type kv struct {
		key   int
		value float64
	}

	// Convert map to a slice of kv type.
	var mapAsSlice []kv
	for k, v := range unsortedMap {
		mapAsSlice = append(mapAsSlice, kv{key: k, value: v})
	}

	sort.Slice(mapAsSlice, func(i, j int) bool {
		return mapAsSlice[i].value < mapAsSlice[j].value
	})

	for i := range mapAsSlice {
		sortOrder = append(sortOrder, mapAsSlice[i].key)
	}
	return
}

// Optimize optimizes the leaf ordering of a dendrogram using the method
// of Bar-Joseph, et al. 2001.
func Optimize(dendrogram []typedef.SubCluster, dist [][]float64) (optimized []typedef.SubCluster) {

	// Number of nodes.
	n := len(dendrogram)

	// Get leafs beneath each node and group them into two pools: leafs on the left (a)
	// go into one slice and leafs on the right (b) go into a second slice.
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

	// Initialize score map and set zero values for leafs. This is a 3D map with
	// the first dimension corresponding a node and the second and third
	// dimensions corresponding to leaf pairs. The 2D leaf will be the left most
	// leaf of a pair and the 3D leaf will be its rightmost pair. The float64
	// value is the between-leaf distance for that pair.
	m := make(map[int]map[int]map[int]float64, 2*n+1)
	for i := 0; i <= n; i++ {
		m[i] = make(map[int]map[int]float64, 1)
		m[i][i] = make(map[int]float64, 1)
		m[i][i][i] = 0
	}

	// Calculate optimal ordering score for each node.
	for _, cluster := range dendrogram {
		node := cluster.Node
		numLeafsA := len(nodeLeafs[node].a)
		numLeafsB := len(nodeLeafs[node].b)

		// Initialize 2D and 3D maps.
		m[node] = make(map[int]map[int]float64, numLeafsA+numLeafsB)
		for _, leaf := range nodeLeafs[node].a {
			m[node][leaf] = make(map[int]float64, numLeafsB)
		}
		for _, leaf := range nodeLeafs[node].b {
			m[node][leaf] = make(map[int]float64, numLeafsA)
		}

		// Iterate over leafs in pool a and b and generate scores.
		for _, aLeaf := range nodeLeafs[node].a {

			// Sort left nodes scores.
			aSortOrder := sortMap(m[cluster.Leafa][aLeaf])
			for _, bLeaf := range nodeLeafs[node].b {

				// Calculate min distance between all rightmost leafs of aLeaf and
				// leftmost leafs of bLeaf.
				minDist := math.MaxFloat32
				for rightALeaf := range m[cluster.Leafa][aLeaf] {
					for leftBLeaf := range m[cluster.Leafb][bLeaf] {
						if dist[rightALeaf][leftBLeaf] < minDist {
							minDist = dist[rightALeaf][leftBLeaf]
						}
					}
				}

				// Sort right nodes scores.
				bSortOrder := sortMap(m[cluster.Leafb][bLeaf])

				// Calculate score for current node.
				optScore := optimal(aSortOrder, bSortOrder, minDist, m[cluster.Leafa][aLeaf], m[cluster.Leafb][bLeaf], dist)
				m[node][aLeaf][bLeaf] = optScore
				m[node][bLeaf][aLeaf] = optScore
			}
		}
	}

	// Re-order dendrogram.
	optimized = make([]typedef.SubCluster, n)
	copy(optimized, dendrogram)

	// Constraints contains the left and right contraints for each node. -1 is used
	// to indicate there is no constraint.
	constrain := make(map[int]constraints, n)
	constrain[dendrogram[n-1].Node] = constraints{left: -1, right: -1}

	// Iterate over nodes and reorder as needed.
	for i := n - 1; i >= 0; i-- {
		node := dendrogram[i].Node

		// Find best leaf pair.
		minDiff := math.MaxFloat64
		var outerA, outerB int
		if constrain[node].left >= 0 {
			for leafb, value := range m[node][constrain[node].left] {
				if value < minDiff {
					minDiff = value
					outerB = leafb
				}
			}
			outerA = constrain[node].left
		} else if constrain[node].right >= 0 {
			for leafa, value := range m[node][constrain[node].right] {
				if value < minDiff {
					minDiff = value
					outerA = leafa
				}
			}
			outerB = constrain[node].right
		} else { // For top node.
			for leafa := range m[node] {
				for leafb, value := range m[node][leafa] {
					if value < minDiff {
						minDiff = value
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
			optimized[i] = typedef.SubCluster{
				Leafa:   dendrogram[i].Leafb,
				Leafb:   dendrogram[i].Leafa,
				Lengtha: dendrogram[i].Lengthb,
				Lengthb: dendrogram[i].Lengtha,
				Node:    dendrogram[i].Node,
			}
		} else {
			optimized[i] = typedef.SubCluster{
				Leafa:   dendrogram[i].Leafa,
				Leafb:   dendrogram[i].Leafb,
				Lengtha: dendrogram[i].Lengtha,
				Lengthb: dendrogram[i].Lengthb,
				Node:    dendrogram[i].Node,
			}
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
