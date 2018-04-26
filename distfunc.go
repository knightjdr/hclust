package hclust

import (
	"errors"
	"math"
)

// DistFunc returns a function for calculating the distance between two vectors.
// Any entries that are zero in both vectors are ignored and vectors must be equal
// length. Default metrix is euclidean.
func DistFunc(metric string) func(x []float64, y []float64) (dist float64, err error) {
	if metric == "binary" {
		// Binary considers two non-zero values to be equivalent.
		binary := func(x []float64, y []float64) (dist float64, err error) {
			if len(x) != len(y) {
				err = errors.New("Vectors for calculating distance must have equal length")
				return
			}
			denominator := float64(0)
			numerator := float64(0)
			for i := range x {
				if x[i] > 0 && y[i] > 0 {
					numerator++
				}
				// Ignore i when both x[i] and y[i] are zero.
				if x[i] > 0 || y[i] > 0 {
					denominator++
				}
			}
			dist = 1 - (numerator / denominator)
			return
		}
		return binary
	} else if metric == "canberra" {
		// Canberra is a weighted version of manhattan.
		canberra := func(x []float64, y []float64) (dist float64, err error) {
			if len(x) != len(y) {
				err = errors.New("Vectors for calculating distance must have equal length")
				return
			}
			dist = 0
			for i := range x {
				// Ignore i when both x[i] and y[i] are zero.
				if x[i] > 0 || y[i] > 0 {
					dist += math.Abs(x[i]-y[i]) / math.Abs(x[i]+y[i])
				}
			}
			return
		}
		return canberra
	} else if metric == "jaccard" {
		// Generalized Jaccard distance.
		jaccard := func(x []float64, y []float64) (dist float64, err error) {
			if len(x) != len(y) {
				err = errors.New("Vectors for calculating distance must have equal length")
				return
			}
			denominator := float64(0)
			numerator := float64(0)
			for i := range x {
				// Ignore i when both x[i] and y[i] are zero.
				if x[i] > 0 || y[i] > 0 {
					numerator += math.Min(x[i], y[i])
					denominator += math.Max(x[i], y[i])
				}
			}
			dist = 1 - (numerator / denominator)
			return
		}
		return jaccard
	} else if metric == "manhattan" {
		// Manhattan sums the differences.
		manhattan := func(x []float64, y []float64) (dist float64, err error) {
			if len(x) != len(y) {
				err = errors.New("Vectors for calculating distance must have equal length")
				return
			}
			dist = 0
			for i := range x {
				dist += math.Abs(x[i] - y[i])
			}
			return
		}
		return manhattan
	} else if metric == "maximum" {
		// Distance between vectors is the maximum difference between elements.
		maximum := func(x []float64, y []float64) (dist float64, err error) {
			if len(x) != len(y) {
				err = errors.New("Vectors for calculating distance must have equal length")
				return
			}
			dist = 0
			for i := range x {
				diff := math.Abs(x[i] - y[i])
				if diff > dist {
					dist = diff
				}
			}
			return
		}
		return maximum
	}
	// Euclidean by default.
	euclidean := func(x []float64, y []float64) (dist float64, err error) {
		if len(x) != len(y) {
			err = errors.New("Vectors for calculating distance must have equal length")
			return
		}
		dist = 0
		for i := range x {
			dist += math.Pow(x[i]-y[i], 2)
		}
		dist = math.Sqrt(dist)
		return
	}
	return euclidean
}
