package hclust

import (
	"math"
	"sort"
	"strconv"
)

// Single clusters a distance matrix using the single (minimum) linkage method.
func Single(matrix [][]float64) (dendrogram []SubCluster) {
	// Number of leafs.
	N := len(matrix)

	// Leaf labels.
	iterLabel := make([]int, N-1)

	// Distance between leafs.
	distance := make(map[string]float64, N-1)

	// Current node.
	c := 0

	// Zero labels and find min distance for first node distance.
	for i := 0; i < N-1; i++ {
		iterLabel[i] = i + 1
		strLabel := strconv.Itoa(i + 1)
		distance[strLabel] = matrix[c][i+1]
	}

	// Iterate until there is a single cluster remaining.
	for i := 0; i < N-1; i++ {
		// Nearest node.
		nodeIndex := ArgMinSingle(distance)
		numIndex, _ := strconv.Atoi(nodeIndex)
		dendrogram = append(
			dendrogram,
			SubCluster{c, numIndex, distance[nodeIndex], distance[nodeIndex]},
		)

		// Change current node.
		c = numIndex

		// Remove new current node from iterLabels.
		cIndex := SliceIndex(len(iterLabel), func(j int) bool { return iterLabel[j] == c })
		iterLabel = append(iterLabel[:cIndex], iterLabel[cIndex+1:]...)

		// Update previous distance.
		distanceLast := distance
		distance = make(map[string]float64, len(iterLabel))
		for _, label := range iterLabel {
			strLabel := strconv.Itoa(label)
			distance[strLabel] = math.Min(distanceLast[strLabel], matrix[label][c])
		}
	}

	// Sort dendrogram.
	sort.SliceStable(dendrogram, func(i, j int) bool {
		return dendrogram[i].Lengtha < dendrogram[j].Lengtha
	})

	// Label dendrogram.
	dendrogram = AddNodes(dendrogram)

	return
}
