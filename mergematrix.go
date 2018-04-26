package hclust

import "math"

// MergeRows merges a pair of rows/columns (a, b) in a symmetric matrix based on the
// specified linkage method.
func MergeRows(method string) func(matrix [][]float64, a, b int) (merged [][]float64) {
	//  Complete or maximum distance.
	complete := func(matrix [][]float64, a, b int) (merged [][]float64) {
		matrix, clusterA, clusterB := RemoveCluster(matrix, a, b)
		// Update first cluster based on linkage method.
		for i := range matrix[a] {
			matrix[a][i] = math.Max(clusterA[i], clusterB[i])
		}
		// Change self distance to 0.
		matrix[a][a] = 0
		return
	}
	return complete
}
