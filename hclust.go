// Package hclust contains methods for performing agglomerative hierarchical clustering.
package hclust

import (
	"github.com/knightjdr/hclust/cluster"
	"github.com/knightjdr/hclust/distance"
	"github.com/knightjdr/hclust/optimize"
	"github.com/knightjdr/hclust/tree"
)

// Cluster references the main cluster method in the cluster subpackage.
var Cluster = cluster.Cluster

// Distance references the main distance method in the distance subpackage.
var Distance = distance.Distance

// Optimize references the main leaf optimization method in the optimize subpackage.
var Optimize = optimize.Optimize

// Tree references the main method for generating the newick tree in the tree subpackage.
var Tree = tree.Create
