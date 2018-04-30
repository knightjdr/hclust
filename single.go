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
	iterLabel := make([]int, N)

	// Distance between leafs.
	distance := make(map[string]float64, N)

	// Zero labels and distance.
	for i := 0; i < N; i++ {
		iterLabel[i] = i
		strLabel := strconv.Itoa(i)
		distance[strLabel] = math.MaxFloat64
	}

	// Current node.
	c := iterLabel[0]

	// Iterate until there is a single cluster remaining.
	for i := 1; i < N; i++ {
		// Remove current node from iterLabels.
		cIndex := SliceIndex(len(iterLabel), func(j int) bool { return iterLabel[j] == c })
		iterLabelNext := make([]int, len(iterLabel)-1)
		if cIndex >= 0 {
			iterLabelNext = append(iterLabel[:cIndex], iterLabel[cIndex+1:]...)
		}

		// Find leaf nearest to current node.
		distanceNext := make(map[string]float64, len(iterLabelNext))
		for _, label := range iterLabelNext {
			strLabel := strconv.Itoa(label)
			distanceNext[strLabel] = math.Min(distance[strLabel], matrix[label][c])
		}

		// Nearest node.
		nodeIndex := ArgMinSingle(distanceNext)
		numIndex, _ := strconv.Atoi(nodeIndex)
		dendrogram = append(
			dendrogram,
			SubCluster{c, numIndex, distanceNext[nodeIndex], distanceNext[nodeIndex]},
		)

		// Change current node.
		c = numIndex

		// Update previous distance and labels.
		distance = distanceNext
		iterLabel = iterLabelNext
	}

	// Sort dendrogram.
	sort.SliceStable(dendrogram, func(i, j int) bool {
		return dendrogram[i].Lengtha < dendrogram[j].Lengtha
	})

	// Label dendrogram.
	dendrogram = AddNodes(dendrogram)

	return
}
