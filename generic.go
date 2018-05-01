package hclust

import (
	"math"
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
	// Update method.
	updateFunc, err := UpdateGeneric(method)
	if err != nil {
		return
	}

	// Number of leafs.
	N := len(matrix)

	// Square values in matrix
	squared := Square(matrix)

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

	// Generate queue with nearest neighbor list.
	queue := make([]NeighborInfo, N)
	for i := 0; i < N-1; i++ {
		neighbor := ArgMinGeneric(squared[i], i)
		queue[i] = NeighborInfo{squared[i][neighbor], i, neighbor}
	}
	queue[N-1] = NeighborInfo{math.MaxFloat64, N - 1, N - 1}

	// Sort queue.
	sort.SliceStable(queue, func(i, j int) bool {
		return queue[i].Dist < queue[j].Dist
	})

	// Iterate over Queue.
	node := N // First node to add.
	for i := 0; i < N-1; i++ {
		// Get element and it's neigbor with shortest distance.
		a := queue[0].Index
		b := queue[0].Neighbor
		delta := queue[0].Dist
		for delta != squared[a][b] {
			neighbor := ArgMinGeneric(squared[a], a)
			queue[0] = NeighborInfo{squared[a][neighbor], a, neighbor}
			// Re-sort queue if a is no longer part of tighest cluster.
			if len(queue) > 1 && queue[0].Dist > queue[1].Dist {
				sort.SliceStable(queue, func(j, k int) bool {
					return queue[j].Dist < queue[k].Dist
				})
				a = queue[0].Index
			}
			b = queue[0].Neighbor
			delta = queue[0].Dist
		}

		// Add new subcluster to dendrogram.
		dendrogram = append(dendrogram, SubCluster{a, b, squared[a][b], squared[a][b]})

		// Remove "a"  and "b" from queue
		queue = queue[1:]
		bIndex := SliceIndex(len(queue), func(j int) bool { return queue[j].Index == b })
		queue = append(queue[:bIndex], queue[bIndex+1:]...)

		// Remove "a" and "b" from labels
		aIndex := SliceIndex(len(labels), func(j int) bool { return labels[j] == a })
		labels = append(labels[:aIndex], labels[aIndex+1:]...)
		bIndex = SliceIndex(len(labels), func(j int) bool { return labels[j] == b })
		labels = append(labels[:bIndex], labels[bIndex+1:]...)

		// Create new node.
		nodeSize[node] = nodeSize[a] + nodeSize[b]
		labels = append(labels, node)

		// Update distance matrix with new node.
		squared = append(squared, updateFunc(squared, a, b, nodeSize)) // Add row.
		for j := 0; j < node; j++ {
			// Add new column.
			squared[j] = append(squared[j], squared[node][j])
			// Set any current distances to A and B to max to exclude them from now on.
			squared[j][a] = math.MaxFloat64
			squared[j][b] = math.MaxFloat64
			squared[node][a] = math.MaxFloat64
			squared[node][b] = math.MaxFloat64
		}

		// Update neighbor candidates that used to be a or b to new node.
		for j := range queue {
			if queue[j].Index < a && queue[j].Neighbor == a {
				queue[j].Neighbor = node
			} else if queue[j].Index < b && queue[j].Neighbor == b {
				queue[j].Neighbor = node
			}
		}

		// If b isn't the previously added node, make the newest node the previous last node's
		// best match.
		if b != node-1 {
			previousIndex := SliceIndex(len(queue), func(j int) bool { return queue[j].Index == node-1 })
			previousNode := queue[previousIndex]
			queue[previousIndex] = NeighborInfo{squared[previousNode.Index][node], previousNode.Index, node}
		}
		// Add the new node to the queue. Reference itself as its best match with
		// infinite distance.
		queue = append(queue, NeighborInfo{math.MaxFloat64, node, node})

		// Re-sort queue.
		sort.SliceStable(queue, func(j, k int) bool {
			return queue[j].Dist < queue[k].Dist
		})

		// Increment node.
		node++
	}

	// Take the square root of all lengths.
	for i := range dendrogram {
		dendrogram[i].Lengtha = math.Sqrt(dendrogram[i].Lengtha)
		dendrogram[i].Lengthb = math.Sqrt(dendrogram[i].Lengthb)
	}

	// Sort dendrogram.
	sort.SliceStable(dendrogram, func(i, j int) bool {
		return dendrogram[i].Lengtha < dendrogram[j].Lengtha
	})

	// Label dendrogram and add branch lengths.
	dendrogram = AddNodes(dendrogram)

	return
}
