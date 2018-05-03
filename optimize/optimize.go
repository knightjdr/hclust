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

func optimal(nodea, nodeb, leafa, leafb int, aSortOrder, bSortOrder []int, minDist float64, nodeScores map[int]map[int]map[int]float64, dist [][]float64) (score float64) {
	// Current best maximal score.
	score = math.MaxFloat64
	for _, leftIndex := range aSortOrder {
		ma := nodeScores[nodea][leafa][leftIndex]
		if ma+nodeScores[nodeb][leafb][bSortOrder[0]]+minDist >= score {
			return
		}
		for _, rightIndex := range bSortOrder {
			mb := nodeScores[nodeb][leafb][rightIndex]
			if ma+mb+minDist >= score {
				break
			} else if score > ma+mb+dist[leftIndex][rightIndex] {
				score = ma + mb + dist[leftIndex][rightIndex]
			}
		}
	}
	return
}

func sortMap(mapArray map[int]float64) (sortOrder []int) {
	type kv struct {
		key   int
		value float64
	}

	var mapObject []kv
	for k, v := range mapArray {
		mapObject = append(mapObject, kv{key: k, value: v})
	}

	sort.Slice(mapObject, func(i, j int) bool {
		return mapObject[i].value < mapObject[j].value
	})

	for i := range mapObject {
		sortOrder = append(sortOrder, mapObject[i].key)
	}
	return
}

// Optimize optimizes the leafs ordering of a dendrogram using the method
// of Bar-Joseph, et al. 2001.
func Optimize(dendrogram []typedef.SubCluster, dist [][]float64) (optimized []typedef.SubCluster) {

	// Number of nodes.
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

	// Initialize score map and set zero values for leafs
	m := make(map[int]map[int]map[int]float64, 2*n+1) // Optimal ordering map.
	for i := 0; i <= n; i++ {
		m[i] = make(map[int]map[int]float64, 1)
		m[i][i] = make(map[int]float64, 1)
		m[i][i][i] = 0
	}

	// Calculate optimal ordering score for each node.
	for _, cluster := range dendrogram {
		node := cluster.Node
		leafs := append(nodeLeafs[node].a, nodeLeafs[node].b...)
		totalLeafs := len(leafs)

		// Initialize 2- and 3D map.
		m[node] = make(map[int]map[int]float64, totalLeafs)
		for _, leaf := range leafs {
			m[node][leaf] = make(map[int]float64, totalLeafs)
		}

		// Calculate min distance between all leafs in pool a against those in pool b.
		minDist := float64(0)
		for _, aLeaf := range nodeLeafs[node].a {
			for _, bLeaf := range nodeLeafs[node].b {
				if dist[aLeaf][bLeaf] < minDist {
					minDist = dist[aLeaf][bLeaf]
				}
			}
		}

		// Iterate over leafs in pool a and b and generate scores.
		for _, aLeaf := range nodeLeafs[node].a {

			// Sort left nodes scores.
			aSortOrder := sortMap(m[cluster.Leafa][aLeaf])
			for _, bLeaf := range nodeLeafs[node].b {

				// Sort right nodes scores.
				bSortOrder := sortMap(m[cluster.Leafb][bLeaf])

				// Calculate score for current node.
				optScore := optimal(cluster.Leafa, cluster.Leafb, aLeaf, bLeaf, aSortOrder, bSortOrder, minDist, m, dist)
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
		} else {
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
