package hclust

// Linkage returns the linkage method to use for finding the next cluster.
// It returns the distance between the pair that will be clustered and the
// indices of the pair.
func Linkage(method string) func(matrix [][]float64) (dist float64, cluster map[string]int) {
	//  Complete or maximum distance.
	complete := func(matrix [][]float64) (dist float64, cluster map[string]int) {
		dim := len(matrix)
		dist = 0
		cluster["a"] = 0
		cluster["b"] = 0
		for i := range matrix {
			for j := i + 1; j < dim; j++ {
				if matrix[i][j] > dist {
					dist = matrix[i][j]
					cluster["a"] = i
					cluster["b"] = j
				}
			}
		}
		return
	}
	return complete
}
