package hclust

// RemoveCluster removes a cluster (row and column) from a symmetric matrix and
// returns the new matrix and vectors containing the removed cluster and the
// cluster it will be merged with.
func RemoveCluster(matrix [][]float64, a, b int) (newMatrix [][]float64, clusterA, clusterB []float64) {
	// Get row for first cluster without the second cluster.
	clusterA = matrix[a]
	clusterA = append(clusterA[:b], clusterA[b+1:]...)
	// Get row for second cluster without itself.
	clusterB = matrix[b]
	clusterB = append(clusterB[:b], clusterB[b+1:]...)

	// Remove second cluster from matrix.
	newMatrix = matrix
	newMatrix = append(newMatrix[:b], newMatrix[b+1:]...) // Remove as row.
	// Remove as column.
	for _, row := range matrix {
		row = append(row[:b], row[b+1:]...)
	}
	return
}
