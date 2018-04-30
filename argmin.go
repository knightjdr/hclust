package hclust

import (
	"math"
)

// ArgMinNN finds the nearest neighbour (with the smallest distance)
// for a node using it's row from a distance matrix and only considering nodes
// greater than it.
func ArgMinGeneric(anchorDist []float64, anchor int, exclude []int) (nearest int) {
	dim := len(anchorDist)
	dist := math.MaxFloat64
	// Create map for exclude slice.
	excludeMap := make(map[int]bool, len(exclude))
	for i := range exclude {
		excludeMap[exclude[i]] = true
	}
	// Find lowest value in anchorDist slice not equal to anchor or in exclude.
	for i := anchor + 1; i < dim; i++ {
		if _, ok := excludeMap[i]; ok {
			continue
		}
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
func ArgMinNN(anchorDist []float64, anchor, preference int, exclude []int) (nearest int) {
	dist := math.MaxFloat64
	// Create map for exclude slice.
	excludeMap := make(map[int]bool, len(exclude))
	for i := range exclude {
		excludeMap[exclude[i]] = true
	}
	// Find lowest value in anchorDist slice not equal to anchor or in exclude.
	for i := range anchorDist {
		if i == anchor {
			continue
		} else if _, ok := excludeMap[i]; ok {
			continue
		}
		if anchorDist[i] < dist {
			dist = anchorDist[i]
			nearest = i
		} else if anchorDist[i] == dist && i == preference {
			dist = anchorDist[i]
			nearest = i
		}
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
