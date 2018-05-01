package hclust

import (
	"math"
)

// ArgMinNN finds the nearest neighbour (with the smallest distance)
// for a node using it's row from a distance matrix and only considering nodes
// greater than it.
func ArgMinGeneric(anchorDist []float64, anchor int) (nearest int) {
	dim := len(anchorDist)
	dist := math.MaxFloat64
	// Find lowest value in anchorDist slice not equal to anchor. Do this in two
	// steps to skip an if statement that excludes anchor.
	for i := anchor + 1; i < dim; i++ {
		if anchorDist[i] < dist {
			dist = anchorDist[i]
			nearest = i
		}
	}
	return
}

// ArgMinNN finds the nearest neighbour (with the smallest distance)
// for a node using it's row from a distance matrix with a preference for another
// node if specified and excluding anything in "exclude".
func ArgMinNN(anchorDist []float64, anchor, preference int) (nearest int) {
	dim := len(anchorDist)
	dist := math.MaxFloat64
	// Find lowest value in anchorDist slice not equal to anchor. Do this in two
	// steps to skip an if statement that excludes anchor.
	for i := 0; i < anchor; i++ {
		if anchorDist[i] < dist {
			dist = anchorDist[i]
			nearest = i
		}
	}
	for i := anchor + 1; i < dim; i++ {
		if anchorDist[i] < dist {
			dist = anchorDist[i]
			nearest = i
		}
	}
	// Check if prefered index has a distance matching lowest distance. iI so, use it.
	if preference >= 0 && anchorDist[preference] == dist {
		nearest = preference
	}
	return
}

// ArgMinSingle finds the nearest neighbour (with the smallest distance)
// for a node using a distance vector. This is to use with the single linkage
// method.
func ArgMinSingle(distVect map[string]float64) (nearest string) {
	dist := math.MaxFloat64
	for key, value := range distVect {
		if value < dist {
			dist = value
			nearest = key
		}
	}
	return
}
