package cluster

import (
	"math"
	"sort"
	"strconv"

	"github.com/knightjdr/hclust/matrixop"
	"github.com/knightjdr/hclust/tree"
	"github.com/knightjdr/hclust/typedef"
)

// Single clusters a distance matrix using the single (minimum) linkage method.
func Single(matrix [][]float64) (dendrogram []typedef.SubCluster) {
	// Number of leafs.
	n := len(matrix)

	// Leaf labels.
	iterLabel := make([]int, n-1)

	// Distance between leafs.
	distance := make(map[string]float64, n-1)

	// Current node.
	c := 0

	// Zero labels and find min distances for first node.
	for i := 0; i < n-1; i++ {
		iterLabel[i] = i + 1
		strLabel := strconv.Itoa(i + 1)
		distance[strLabel] = matrix[c][i+1]
	}

	// Iterate until there is a single cluster remaining.
	for i := 0; i < n-1; i++ {
		// Nearest node.
		nodeIndex := ArgMinSingle(distance)
		numIndex, _ := strconv.Atoi(nodeIndex)
		dendrogram = append(
			dendrogram,
			typedef.SubCluster{
				Leafa:   c,
				Leafb:   numIndex,
				Lengtha: distance[nodeIndex],
				Lengthb: distance[nodeIndex],
				Node:    0,
			},
		)

		// Change current node.
		c = numIndex

		// Remove new current node from iterLabels.
		cIndex := matrixop.SliceIndex(len(iterLabel), func(j int) bool { return iterLabel[j] == c })
		iterLabel = append(iterLabel[:cIndex], iterLabel[cIndex+1:]...)

		// Update previous distance.
		distanceLast := distance
		distance = make(map[string]float64, len(iterLabel))
		for _, label := range iterLabel {
			strLabel := strconv.Itoa(label)
			distance[strLabel] = math.Min(distanceLast[strLabel], matrix[c][label])
		}
	}

	// Sort dendrogram.
	sort.SliceStable(dendrogram, func(i, j int) bool {
		return dendrogram[i].Lengtha < dendrogram[j].Lengtha
	})

	// Label dendrogram.
	dendrogram = tree.AddNodes(dendrogram)

	return
}
