// Package sort will sort a matrix by row or column.
package sort

import (
	"errors"

	"github.com/knightjdr/hclust/matrixop"
)

// Sort takes a 2D matrix and sorts based on the columns or rows. A vector of
// names must be supplied, along with the sorted order of the names. "dim"
// must be one of "column" or "row"
func Sort(matrix [][]float64, names, sortOrder []string, dim string) (sorted [][]float64, err error) {
	sorted = make([][]float64, len(matrix))
	// Ensure the names and sortOrder are of equal length and that they match the
	// length of the dimension to be sorted.
	if len(names) != len(sortOrder) {
		err = errors.New("The vector of unsorted and sorted names must have the same length")
		return
	} else if dim == "column" && len(names) != len(matrix[0]) {
		err = errors.New("The vector of names must be the same length as the dimension to sort")
		return
	} else if dim == "row" && len(names) != len(matrix) {
		err = errors.New("The vector of names must be the same length as the dimension to sort")
		return
	}

	// Create sort map.
	sortedPos := make([]int, len(names))
	for i, name := range names {
		pos := matrixop.SliceIndex(len(sortOrder), func(j int) bool { return sortOrder[j] == name })
		if pos < 0 {
			err = errors.New("Name could not be found in sorted vector")
			return
		}
		sortedPos[i] = pos
	}

	// Sort by column.
	numCols := len(matrix[0])
	if dim == "column" {
		for i, row := range matrix {
			sorted[i] = make([]float64, numCols)
			for j, column := range row {
				sorted[i][sortedPos[j]] = column
			}
		}
	} else { // Sort by row.
		for i, row := range matrix {
			sorted[sortedPos[i]] = row
		}
	}

	return
}
