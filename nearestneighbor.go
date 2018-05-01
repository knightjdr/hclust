package hclust

import (
	"math"
	"sort"
)

// NearestNeighbor clusters a distance matrix using one of the following linkage
// methods: average, complete, mcquitty or ward.
func NearestNeighbor(matrix [][]float64, method string) (dendrogram []SubCluster, err error) {
	// Number of leafs.
	N := len(matrix)

	// Square the matrix for wards
	dist := matrix
	if method == "ward" {
		dist = Square(matrix)
	}

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
	for len(labels) > 1 {
		var a, b int

		// Get nodes to test for as neighbors.
		a = labels[0] // Grab any node (use first).
		chain = append(chain, a)
		b = labels[1] // Grab any node besides a (grab second)

		// Find nearest neighbor of node a.
		for len(chain) < 3 || (a != chain[len(chain)-3]) {
			c := ArgMinNN(dist[a], a, b)
			b = a
			a = c
			chain = append(chain, a)
		}

		// Add new cluster to dendrogram.
		dendrogram = append(dendrogram, SubCluster{a, b, dist[a][b], dist[a][b]})

		// Remove a and b from labels.
		aIndex := SliceIndex(len(labels), func(i int) bool { return labels[i] == a })
		labels = append(labels[:aIndex], labels[aIndex+1:]...)
		bIndex := SliceIndex(len(labels), func(i int) bool { return labels[i] == b })
		labels = append(labels[:bIndex], labels[bIndex+1:]...)

		// New node.
		nodeSize[node] = nodeSize[a] + nodeSize[b]

		// Update distance matrix with new node.
		dist = append(dist, updateFunc(dist, a, b, nodeSize)) // Add new row.
		for i := 0; i < node; i++ {
			// Add new column.
			dist[i] = append(dist[i], dist[node][i])
			// Set any current distances to A and B to max to exclude them from now on.
			dist[i][a] = math.MaxFloat64
			dist[i][b] = math.MaxFloat64
			dist[node][a] = math.MaxFloat64
			dist[node][b] = math.MaxFloat64
		}

		// Append node.
		labels = append(labels, node)

		// Increment node.
		node++
	}

	// Take the square root of all lengths for ward.
	if method == "ward" {
		for i := range dendrogram {
			dendrogram[i].Lengtha = math.Sqrt(dendrogram[i].Lengtha)
			dendrogram[i].Lengthb = math.Sqrt(dendrogram[i].Lengthb)
		}
	}

	// Sort dendrogram.
	sort.SliceStable(dendrogram, func(i, j int) bool {
		return dendrogram[i].Lengtha < dendrogram[j].Lengtha
	})

	// Label dendrogram and add branch lengths.
	dendrogram = AddNodes(dendrogram)

	return
}
