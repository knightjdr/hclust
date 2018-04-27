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

	// Leaf labels for loop iterations.
	iterLabels := make([][]int, N)
	iterLabels[0] = make([]int, N)
	for i := range iterLabels[0] {
		iterLabels[0][i] = i
	}

	// Distance between leafs.
	distance := make([]map[string]float64, N)
	distance[0] = make(map[string]float64, N)
	for i := range iterLabels[0] {
		strLabel := strconv.Itoa(i)
		distance[0][strLabel] = math.MaxFloat64
	}

	// Current node.
	c := iterLabels[0][0]

	// Iterate until there is a single cluster remaining.
	for i := 1; i < N; i++ {
		// Remove current node from iterLabels.
		cIndex := SliceIndex(len(iterLabels[i-1]), func(j int) bool { return iterLabels[i-1][j] == c })
		iterLabels[i] = make([]int, len(iterLabels[i-1])-1)
		if cIndex >= 0 {
			iterLabels[i] = append(iterLabels[i-1][:cIndex], iterLabels[i-1][cIndex+1:]...)
		}

		// Find leaf nearest to current node.
		distance[i] = make(map[string]float64, len(iterLabels[i]))
		for _, label := range iterLabels[i] {
			strLabel := strconv.Itoa(label)
			distance[i][strLabel] = math.Min(distance[i-1][strLabel], matrix[label][c])
		}

		// Nearest node.
		nodeIndex := ArgMinSingle(distance[i])
		numIndex, _ := strconv.Atoi(nodeIndex)
		dendrogram = append(
			dendrogram,
			SubCluster{distance[i][nodeIndex], c, numIndex},
		)

		// Change current node.
		c = numIndex
	}

	// Sort dendrogram.
	sort.SliceStable(dendrogram, func(i, j int) bool {
		return dendrogram[i].Dist < dendrogram[j].Dist
	})

	// Label dendrogram.
	dendrogram = AddNodes(dendrogram)

	return
}
