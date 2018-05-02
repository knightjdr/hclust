package matrixop

// Square squares every element in a matrix.
func Square(matrix [][]float64) (squared [][]float64) {
	m := len(matrix)
	n := len(matrix[0])
	squared = make([][]float64, m)
	for i := range matrix {
		squared[i] = make([]float64, n)
		for j := range matrix[i] {
			squared[i][j] = matrix[i][j] * matrix[i][j]
		}
	}
	return
}
