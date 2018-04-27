package hclust

import "math"

// ArgMin finds the nearest neighbour (with the smallest distance)
// for a node using it's row from a distance matrix with a preference for another
// node if specified.
func ArgMin(anchorDist []float64, anchor, preference int) (nearest int) {
	dist := math.MaxFloat64
	for i := range anchorDist {
		if i == anchor {
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
