package hclust

import "errors"

// SubCluster stores the distance and names of leafs for a subcluster.
type SubCluster struct {
	dist  float64
	leafl interface{}
	leafr interface{}
}

// Cluster clusters a square symmetric matrix and returns a dendrogram as a
// string as well as an ordered vector matching the dendrogram. A vector with the
// row/column names is required for the dendrogram and ordered vector. Linkage
// method options are: average, centroid, complete, McQuitty,
// median, single and Ward’s.
func Cluster(matrix [][]float64, names []string, method string) (dendrogram string, ordered []string, err error) {
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

	// Linkage functions.
	linkage := Linkage(method)
	mergeRows := MergeRows(method)

	// Create cluster list for holding cluster leafs and subcluster names.
	// Leafs are denoted by strings and subclusters by integers.
	clusterList := make([]interface{}, len(names))
	for i := range names {
		clusterList[i] = names[i]
	}

	// Iterate until cluster number equals N - 1
	clusteredMatrix := matrix           // Create copy of input matrix.
	clusters := make([]SubCluster, N-1) // Store each cluster by index.
	for clusterID := 0; clusterID < N-1; clusterID++ {
		// Get most similar cluster.
		dist, subCluster := linkage(clusteredMatrix)
		// Create new cluster and add to cluster names.
		clusters[clusterID], clusterList = CreateCluster(clusterList, clusterID, dist, subCluster["a"], subCluster["b"])
		// Update matrix.
		clusteredMatrix = mergeRows(clusteredMatrix, subCluster["a"], subCluster["b"])
	}
	return
}
