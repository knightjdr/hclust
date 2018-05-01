package hclust

// Square squares every element in a matrix.
func Square(matrix [][]float64) (squared [][]float64) {
	M := len(matrix)
	N := len(matrix[0])
	squared = make([][]float64, M)
	for i := range matrix {
		squared[i] = make([]float64, N)
		for j := range matrix[i] {
			squared[i][j] = matrix[i][j] * matrix[i][j]
		}
	}
	return
}
