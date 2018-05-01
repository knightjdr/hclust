package hclust

import "errors"

// SubCluster stores the distance and names of leafs for a subcluster.
type SubCluster struct {
	Leafa   int
	Leafb   int
	Lengtha float64
	Lengthb float64
}

// Cluster clusters a square symmetric matrix and returns a dendrogram as a
// string as well as an ordered vector matching the dendrogram. A vector with the
// row/column names is required for the dendrogram and ordered vector. Linkage
// method options are: average, centroid, complete, McQuitty,
// median, single and Ward’s.
func Cluster(matrix [][]float64, names []string, method string) (dendrogram []SubCluster, order []string, err error) {
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
	dendrogram = make([]SubCluster, 2*N-1)
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

	// Get newick tree and cluster order.
	Tree(dendrogram, names)

	return
}
