// Package cluster contains methods for clustering.
package cluster

import (
	"errors"

	"github.com/knightjdr/hclust/typedef"
)

// Cluster clusters a square symmetric matrix and returns a dendrogram as a
// string as well as an ordered vector matching the dendrogram. A vector with the
// row/column names is required for the dendrogram and ordered vector. Linkage
// method options are: average, centroid, complete, McQuitty,
// median, single and Ward’s.
func Cluster(matrix [][]float64, method string) (dendrogram []typedef.SubCluster, err error) {
	// Return if matrix is not symmetric.
	colDim := len(matrix[0])
	rowDim := len(matrix)
	if colDim != rowDim {
		err = errors.New("The matrix must be symmetric")
		return
	}
	N := rowDim // Matrix dimension.

	// Linkage.
	dendrogram = make([]typedef.SubCluster, 2*N-1)
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
