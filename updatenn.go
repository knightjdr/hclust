package hclust

import (
	"errors"
	"math"
)

// UpdateNN calculates the new row/column to add to a distance matrix for a new node.
// Methods supported: average, complete, mcquitty or ward.
func UpdateNN(method string) (updateFunc func(matrix [][]float64, a, b int, nodeSize []int) (newRow []float64), err error) {
	if method == "average" {
		average := func(matrix [][]float64, a, b int, nodeSize []int) (newRow []float64) {
			x := matrix[a]
			y := matrix[b]
			dim := len(x)
			newRow = make([]float64, dim+1)
			for i := 0; i < dim; i++ {
				numerator := (float64(nodeSize[a]) * x[i]) + (float64(nodeSize[b]) * y[i])
				denominator := float64(nodeSize[a] + nodeSize[b])
				newRow[i] = numerator / denominator
			}
			// Set self distance to zero.
			newRow[dim] = 0
			return
		}
		return average, nil
	} else if method == "complete" {
		complete := func(matrix [][]float64, a, b int, nodeSize []int) (newRow []float64) {
			x := matrix[a]
			y := matrix[b]
			dim := len(x)
			newRow = make([]float64, dim+1)
			for i := 0; i < dim; i++ {
				newRow[i] = math.Max(x[i], y[i])
			}
			// Set self distance to zero.
			newRow[dim] = 0
			return
		}
		return complete, nil
	} else if method == "mcquitty" {
		mcquitty := func(matrix [][]float64, a, b int, nodeSize []int) (newRow []float64) {
			x := matrix[a]
			y := matrix[b]
			dim := len(x)
			newRow = make([]float64, dim+1)
			for i := 0; i < dim; i++ {
				newRow[i] = (x[i] + y[i]) / float64(2)
			}
			// Set self distance to zero.
			newRow[dim] = 0
			return
		}
		return mcquitty, nil
	} else if method == "ward" {
		ward := func(matrix [][]float64, a, b int, nodeSize []int) (newRow []float64) {
			x := matrix[a]
			y := matrix[b]
			dim := len(x)
			newRow = make([]float64, dim+1)
			for i := 0; i < dim; i++ {
				numerator := float64(nodeSize[a]+nodeSize[i]) * math.Pow(x[i], 2)
				numerator += float64(nodeSize[b]+nodeSize[i]) * math.Pow(y[i], 2)
				numerator -= float64(nodeSize[i]) * math.Pow(x[b], 2)
				denominator := float64(nodeSize[a] + nodeSize[b] + nodeSize[i])
				newRow[i] = math.Sqrt(numerator / denominator)
			}
			// Set self distance to zero.
			newRow[dim] = 0
			return
		}
		return ward, nil
	}
	err = errors.New("Unknown linkage method")
	return
}
