package hclust

import (
	"errors"
	"math"
)

// UpdateGeneric calculates the new row/column to add to a distance matrix for a new node.
// Methods supported: centroid or median.
func UpdateGeneric(method string) (updateFunc func(matrix [][]float64, a, b int, nodeSize []int) (newRow []float64), err error) {
	if method == "centroid" {
		centroid := func(matrix [][]float64, a, b int, nodeSize []int) (newRow []float64) {
			x := matrix[a]
			y := matrix[b]
			dim := len(x)
			newRow = make([]float64, dim+1)
			for i := 0; i < dim; i++ {
				leftNumerator := float64(nodeSize[a]) * x[i]
				leftNumerator += float64(nodeSize[b]) * y[i]
				leftDenomimnator := float64(nodeSize[a] + nodeSize[b])
				rightNumerator := float64(nodeSize[a]) * float64(nodeSize[b]) * x[b]
				rightDenomimnator := math.Pow(float64(nodeSize[a]+nodeSize[b]), 2)
				newRow[i] = (leftNumerator / leftDenomimnator) - (rightNumerator / rightDenomimnator)
			}
			// Set self distance to zero.
			newRow[dim] = 0
			return
		}
		return centroid, nil
	}
	if method == "median" {
		median := func(matrix [][]float64, a, b int, nodeSize []int) (newRow []float64) {
			x := matrix[a]
			y := matrix[b]
			dim := len(x)
			newRow = make([]float64, dim+1)
			for i := 0; i < dim; i++ {
				numerator := float64(2) * (x[i] + y[i])

				numerator -= x[b]
				newRow[i] = numerator / float64(4)
			}
			// Set self distance to zero.
			newRow[dim] = 0
			return
		}
		return median, nil
	}
	err = errors.New("Unknown linkage method")
	return
}
