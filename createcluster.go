package hclust

// CreateCluster takes a current subcluster pair (a, b) and creates a subcluster
// of type SubCluster for them. The pair are removed from the current cluster list and a new
// entry is added to the cluster list to denote the subcluster.
func CreateCluster(
	clusterList []interface{},
	clusterID int,
	dist float64,
	a, b int,
) (subcluster SubCluster, updatedList []interface{}) {
	// Create subcluster.
	subcluster = SubCluster{dist, clusterList[a], clusterList[b]}

	// Remove second cluster item from cluster list.
	updatedList = append(clusterList[:b], clusterList[b+1:]...)

	// Change name of first cluster item to cluster ID.
	updatedList[a] = clusterID
	return
}
