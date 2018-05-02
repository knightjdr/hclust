// Package hclust contains methods for performing agglomerative hierarchical clustering.
package hclust

import (
	"github.com/knightjdr/hclust/cluster"
	"github.com/knightjdr/hclust/distance"
)

// Cluster references the main cluster method in the cluster subpackage.
var Cluster = cluster.Cluster

// Distance references the main distance method in the distance subpackage.
var Distance = distance.Distance
