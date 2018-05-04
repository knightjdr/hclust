// Package cluster contains methods for clustering.
package cluster

import (
	"errors"

	"github.com/knightjdr/hclust/typedef"
)

// Cluster clusters a square symmetric matrix and returns a dendrogram. Linkage
// method options are: average, centroid, complete, mcquitty, median, single and ward.
func Cluster(matrix [][]float64, method string) (dendrogram []typedef.SubCluster, err error) {
	// Return if matrix row and column numbers are not equal. This is to ensure
	// the matrix is symmetric (likely will be in this case).
	colDim := len(matrix[0])
	rowDim := len(matrix)
	if colDim != rowDim {
		err = errors.New("The matrix must be symmetric")
		return
	}

	// Matrix dimension.
	N := rowDim

	// Linkage.
	dendrogram = make([]typedef.SubCluster, N-1)
	if method == "single" {
		dendrogram = Single(matrix)
	} else if method == "average" {
		dendrogram, err = NearestNeighbor(matrix, method)
	} else if method == "complete" {
		dendrogram, err = NearestNeighbor(matrix, method)
	} else if method == "mcquitty" {
		dendrogram, err = NearestNeighbor(matrix, method)
	} else if method == "ward" {
		dendrogram, err = NearestNeighbor(matrix, method)
	} else if method == "centroid" {
		dendrogram, err = Generic(matrix, method)
	} else if method == "median" {
		dendrogram, err = Generic(matrix, method)
	} else {
		err = errors.New("Unkown linkage method")
	}

	return
}
