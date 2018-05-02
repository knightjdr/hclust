package matrixop

// SliceIndex finds the index of an element in a slice based on the predicate
// function.
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
