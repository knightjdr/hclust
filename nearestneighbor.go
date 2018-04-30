package hclust

import (
	"sort"
)

// NearestNeighbor clusters a distance matrix using one of the following linkage
// methods: average, complete, mcquitty or ward.
func NearestNeighbor(matrix [][]float64, method string) (dendrogram []SubCluster, err error) {
	// Number of leafs.
	N := len(matrix)

	// Leaf labels.
	labels := make([]int, N)
	for i := 0; i < N; i++ {
		labels[i] = i
	}

	// Number of leafs at each node (including node if it is a leaf).
	nodeSize := make([]int, 2*N-1)
	for i := 0; i < 2*N-1; i++ {
		nodeSize[i] = 1
	}

	// Update method.
	updateFunc, err := UpdateNN(method)
	if err != nil {
		return
	}

	// Iterate until there is a single cluster remaining.
	chain := make([]int, 0)
	node := N // First node to add.
	nodesAdded := make([]int, 0)
	for len(labels) > 1 {
		var a, b int

		// Get nodes to test for as neighbors.
		a = labels[0] // Grab any node (use first).
		chain = append(chain, a)
		b = labels[1] // Grab any node besides a (grab second)

		// Find nearest neighbor of node a.
		for len(chain) < 3 || (a != chain[len(chain)-3]) {
			c := ArgMinNN(matrix[a], a, b, nodesAdded)
			b = a
			a = c
			chain = append(chain, a)
		}

		// Add new cluster to dendrogram.
		dendrogram = append(dendrogram, SubCluster{a, b, matrix[a][b], matrix[a][b]})

		// Remove a and b from labels.
		aIndex := SliceIndex(len(labels), func(i int) bool { return labels[i] == a })
		if aIndex >= 0 {
			labels = append(labels[:aIndex], labels[aIndex+1:]...)
		}
		bIndex := SliceIndex(len(labels), func(i int) bool { return labels[i] == b })
		if bIndex >= 0 {
			labels = append(labels[:bIndex], labels[bIndex+1:]...)
		}

		// New node.
		nodeSize[node] = nodeSize[a] + nodeSize[b]

		// Update distance matrix with new node.
		matrix = append(matrix, updateFunc(matrix, a, b, nodeSize))
		for i := 0; i < node; i++ {
			matrix[i] = append(matrix[i], matrix[node][i])
		}

		// Append node.
		labels = append(labels, node)

		// Increment node and add found nodes to exclude slice.
		nodesAdded = append(nodesAdded, []int{a, b}...)
		node++
	}

	// Sort dendrogram.
	sort.SliceStable(dendrogram, func(i, j int) bool {
		return dendrogram[i].Lengtha < dendrogram[j].Lengtha
	})

	// Label dendrogram and add branch lengths.
	dendrogram = AddNodes(dendrogram)

	return
}
