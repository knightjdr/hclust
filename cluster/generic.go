package cluster

import (
	"math"
	"sort"

	"github.com/knightjdr/hclust/matrixop"
	"github.com/knightjdr/hclust/tree"
	"github.com/knightjdr/hclust/typedef"
)

// NeighborInfo stores information about a nodes nearest neighbor.
type NeighborInfo struct {
	Dist     float64
	Index    int
	Neighbor int
}

// Generic clusters a distance matrix using a generic algorithm and one of the
// following linkage methods: centroid or median.
func Generic(matrix [][]float64, method string) (dendrogram []typedef.SubCluster, err error) {
	// Update method.
	updateFunc, err := UpdateGeneric(method)
	if err != nil {
		return
	}

	// Number of leafs.
	n := len(matrix)

	// Square values in matrix
	dist := matrixop.Square(matrix)

	// Leaf labels.
	labels := make([]int, n)
	for i := 0; i < n; i++ {
		labels[i] = i
	}

	// Number of leafs at each node/leaf.
	nodeSize := make([]int, 2*n-1)
	for i := 0; i < 2*n-1; i++ {
		nodeSize[i] = 1
	}

	// Generate queue with nearest neighbor list.
	queue := make([]NeighborInfo, n)
	for i := 0; i < n-1; i++ {
		neighbor := ArgMinGeneric(dist[i], i)
		queue[i] = NeighborInfo{dist[i][neighbor], i, neighbor}
	}
	// Add last node with itself as nearest neighbor and infinite distance. Need
	// this for code logic below.
	queue[n-1] = NeighborInfo{math.MaxFloat64, n - 1, n - 1}

	// Sort queue.
	sort.SliceStable(queue, func(i, j int) bool {
		return queue[i].Dist < queue[j].Dist
	})

	// Iterate over Queue.
	node := n // First node to add.
	for i := 0; i < n-1; i++ {
		// Get element and its neigbor with shortest distance.
		a := queue[0].Index
		b := queue[0].Neighbor
		delta := queue[0].Dist

		// If b is not a's nearest neigbor, find it. This discrepency happens as
		// nodes get created.
		for delta != dist[a][b] {
			neighbor := ArgMinGeneric(dist[a], a)
			queue[0] = NeighborInfo{dist[a][neighbor], a, neighbor}
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
		dendrogram = append(
			dendrogram,
			typedef.SubCluster{
				Leafa:   a,
				Leafb:   b,
				Lengtha: dist[a][b],
				Lengthb: dist[a][b],
				Node:    node,
			},
		)

		// Remove "a"  and "b" from queue
		queue = queue[1:]
		bIndex := matrixop.SliceIndex(len(queue), func(j int) bool { return queue[j].Index == b })
		queue = append(queue[:bIndex], queue[bIndex+1:]...)

		// Remove "a" and "b" from labels
		aIndex := matrixop.SliceIndex(len(labels), func(j int) bool { return labels[j] == a })
		labels = append(labels[:aIndex], labels[aIndex+1:]...)
		bIndex = matrixop.SliceIndex(len(labels), func(j int) bool { return labels[j] == b })
		labels = append(labels[:bIndex], labels[bIndex+1:]...)

		// Create new node.
		nodeSize[node] = nodeSize[a] + nodeSize[b]
		labels = append(labels, node)

		// Update distance matrix with new node.
		dist = append(dist, updateFunc(dist, a, b, nodeSize)) // Add row.
		for j := 0; j < node; j++ {
			// Add new column.
			dist[j] = append(dist[j], dist[node][j])
			// Set any current distances to a and b to max to exclude them from now on.
			dist[j][a] = math.MaxFloat64
			dist[j][b] = math.MaxFloat64
			dist[node][a] = math.MaxFloat64
			dist[node][b] = math.MaxFloat64
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
			previousIndex := matrixop.SliceIndex(len(queue), func(j int) bool { return queue[j].Index == node-1 })
			queue[previousIndex] = NeighborInfo{dist[node-1][node], node - 1, node}
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
	dendrogram = tree.AddNodes(dendrogram)

	return
}
