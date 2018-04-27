package hclust

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceIndex(t *testing.T) {
	// TEST 1: remove character.
	dataChar := []string{"a", "b", "c"}
	predicate := func(i int) bool { return dataChar[i] == "a" }
	want := 0
	assert.Equal(t, want, SliceIndex(3, predicate), "Character index not found")

	// TEST 2: remove float.
	dataFloat := []float64{1.1, 5.23, 2.01}
	predicate = func(i int) bool { return dataFloat[i] == 2.01 }
	want = 2
	assert.Equal(t, want, SliceIndex(3, predicate), "Float index not found")

	// TEST 3: remove integer.
	dataInt := []int{1, 5, 2}
	predicate = func(i int) bool { return dataInt[i] == 5 }
	want = 1
	assert.Equal(t, want, SliceIndex(3, predicate), "Integer index not found")

	// TEST 4: not found.
	dataInt = []int{1, 5, 2}
	predicate = func(i int) bool { return dataInt[i] == 7 }
	want = -1
	assert.Equal(t, want, SliceIndex(3, predicate), "-1 not returned when value not found")
}
