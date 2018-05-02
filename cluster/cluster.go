// Package cluster contains methods for clustering.
package cluster

import (
	"errors"

	"github.com/knightjdr/hclust/tree"
	"github.com/knightjdr/hclust/typedef"
)

// Cluster clusters a square symmetric matrix and returns a dendrogram as a
// string as well as an ordered vector matching the dendrogram. A vector with the
// row/column names is required for the dendrogram and ordered vector. Linkage
// method options are: average, centroid, complete, McQuitty,
// median, single and Ward’s.
func Cluster(matrix [][]float64, names []string, method string, optimize bool) (clust typedef.Hclust, err error) {
	// Return if matrix is not symmetric.
	colDim := len(matrix[0])
	rowDim := len(matrix)
	if colDim != rowDim {
		err = errors.New("The matrix must be symmetric")
		return
	}
	N := rowDim // Matrix dimension.

	// Return if names length does not match matix length.
	if len(names) != N {
		err = errors.New("The name vector must have the same dimension as the matrix")
		return
	}

	// Linkage.
	clust.Dendrogram = make([]typedef.SubCluster, 2*N-1)
	if method == "single" {
		clust.Dendrogram = Single(matrix)
	} else if method == "average" {
		clust.Dendrogram, err = NearestNeighbor(matrix, method)
	} else if method == "complete" {
		clust.Dendrogram, err = NearestNeighbor(matrix, method)
	} else if method == "mcquitty" {
		clust.Dendrogram, err = NearestNeighbor(matrix, method)
	} else if method == "ward" {
		clust.Dendrogram, err = NearestNeighbor(matrix, method)
	} else if method == "centroid" {
		clust.Dendrogram, err = Generic(matrix, method)
	} else if method == "median" {
		clust.Dendrogram, err = Generic(matrix, method)
	} else {
		err = errors.New("Unkown linkage method")
	}

	// Optimize leaf ordering.
	if optimize {
	}

	// Get newick tree and cluster order.
	newTree := tree.Create(clust.Dendrogram, names)
	clust.Newick = newTree.Newick
	clust.Order = newTree.Order

	return
}
