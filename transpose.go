package hclust

// Transpose transposes a 2D matrix (2D slice)
func Transpose(matrix [][]float64) (transposed [][]float64) {
	// Matrix dimensions.
	colNum := len(matrix[0])
	rowNum := len(matrix)

	// Init transposed matrix.
	transposed = make([][]float64, colNum) // Set row capacity.
	for i := range transposed {
		transposed[i] = make([]float64, rowNum) // Set column capacity.
	}
	for i, row := range matrix {
		for j, value := range row {
			transposed[j][i] = value
		}
	}
	return
}
