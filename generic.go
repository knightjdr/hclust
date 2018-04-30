package hclust

import (
	"sort"
)

// NeighborInfo stores information about a nodes nearest neighbor.
type NeighborInfo struct {
	Dist     float64
	Index    int
	Neighbor int
}

// Generic clusters a distance matrix using a generic algorithm and one of the
// following linkage methods: centroid or median.
func Generic(matrix [][]float64, method string) (dendrogram []SubCluster, err error) {
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
	updateFunc, err := UpdateGeneric(method)
	if err != nil {
		return
	}

	// Generate queue with nearest neighbor list.
	queue := make([]NeighborInfo, N-1)
	for i := 0; i < N-1; i++ {
		neighbor := ArgMinGeneric(matrix[i], i, []int{})
		queue[i] = NeighborInfo{matrix[i][neighbor], i, neighbor}
	}

	// Sort queue.
	sort.SliceStable(queue, func(i, j int) bool {
		return queue[i].Dist < queue[j].Dist
	})

	// Iterate over Queue.
	node := N // First node to add.
	nodesAdded := make([]int, 0)
	for i := 0; i < N-1; i++ {
		// Get element and it's neigbor with shortest distance.
		a := queue[0].Index
		b := queue[0].Neighbor
		z := labels[len(labels)-1]
		delta := queue[0].Dist
		for delta != matrix[a][b] {
			neighbor := ArgMinGeneric(matrix[a], a, nodesAdded)
			queue[0] = NeighborInfo{matrix[a][neighbor], a, neighbor}
			// Resort queue.
			sort.SliceStable(queue, func(j, k int) bool {
				return queue[j].Dist < queue[k].Dist
			})
			a = queue[0].Index
			b = queue[0].Neighbor
			delta = queue[0].Dist
		}

		// Remove "a" from queue
		aIndex := SliceIndex(len(queue), func(j int) bool { return queue[j].Index == a })
		if aIndex >= 0 {
			queue = append(queue[:aIndex], queue[aIndex+1:]...)
		}

		// Remove "a" and "b" from labels
		aIndex = SliceIndex(len(labels), func(i int) bool { return labels[i] == a })
		if aIndex >= 0 {
			labels = append(labels[:aIndex], labels[aIndex+1:]...)
		}
		bIndex := SliceIndex(len(labels), func(i int) bool { return labels[i] == b })
		if bIndex >= 0 {
			labels = append(labels[:bIndex], labels[bIndex+1:]...)
		}

		// Add new subcluster to dendrogram.
		dendrogram = append(dendrogram, SubCluster{a, b, matrix[a][b], matrix[a][b]})

		// Create new node.
		nodeSize[node] = nodeSize[a] + nodeSize[b]
		labels = append(labels, node)

		// Update distance matrix with new node.
		matrix = append(matrix, updateFunc(matrix, a, b, nodeSize))
		for i := 0; i < node; i++ {
			matrix[i] = append(matrix[i], matrix[node][i])
		}

		// Update neighbor candidates that used to be a or b to new node.
		for j := range queue {
			if queue[j].Index < a && queue[j].Neighbor == a {
				queue[j].Neighbor = node
			} else if queue[j].Index < b && queue[j].Neighbor == b {
				queue[j].Neighbor = node
			}
		}

		// Set previous last added node to have new node as nearest neighbor.
		if b != z {
			// Find index of B and update queue with last new node.
			bIndex := SliceIndex(len(queue), func(j int) bool { return queue[j].Index == b })
			queue[bIndex] = NeighborInfo{matrix[node][z], z, node}
			// Resort queue.
			sort.SliceStable(queue, func(j, k int) bool {
				return queue[j].Dist < queue[k].Dist
			})
		}

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
