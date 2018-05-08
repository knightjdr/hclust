package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	// TEST1: names and sorted vector do not match.
	matrix := make([][]float64, 0)
	names := []string{"column1", "column2", "column3"}
	sortOrder := []string{"column2", "column1"}
	want := "The vector of unsorted and sorted names must have the same length"
	_, err := Sort(matrix, names, sortOrder, "column")
	assert.Equal(t, want, err.Error(), "Unsorted and sorted vectors of different length should return correct error")

	// TEST2: names not matching column sort dimension length.
	matrix = [][]float64{
		{1, 2},
		{2, 3},
		{3, 1},
	}
	names = []string{"column1", "column2", "column3"}
	sortOrder = []string{"column2", "column3", "column1"}
	want = "The vector of names must be the same length as the dimension to sort"
	_, err = Sort(matrix, names, sortOrder, "column")
	assert.Equal(t, want, err.Error(), "Unsorted and sorted vectors of different length should return correct error")

	// TEST3: names not matching row sort dimension length.
	matrix = [][]float64{
		{1, 2, 3},
		{2, 3, 1},
	}
	names = []string{"row1", "row2", "row3"}
	sortOrder = []string{"row2", "row3", "row1"}
	want = "The vector of names must be the same length as the dimension to sort"
	_, err = Sort(matrix, names, sortOrder, "row")
	assert.Equal(t, want, err.Error(), "Unsorted and sorted vectors of different length should return correct error")

	// TEST4: name not found in sorted vector.
	matrix = [][]float64{
		{1, 2, 3},
		{2, 3, 1},
		{3, 1, 2},
	}
	names = []string{"column1", "column2", "column4"}
	sortOrder = []string{"column2", "column3", "column1"}
	want = "Name could not be found in sorted vector"
	_, err = Sort(matrix, names, sortOrder, "column")
	assert.Equal(t, want, err.Error(), "Name missing in sorted vector should return correct error")

	// TEST5: sort by column.
	matrix = [][]float64{
		{1, 2, 3},
		{2, 3, 1},
		{3, 1, 2},
	}
	names = []string{"column1", "column2", "column3"}
	sortOrder = []string{"column2", "column3", "column1"}
	wantSorted := [][]float64{
		{2, 3, 1},
		{3, 1, 2},
		{1, 2, 3},
	}
	sorted, _ := Sort(matrix, names, sortOrder, "column")
	assert.Equal(t, wantSorted, sorted, "Matrix not sorted correctly by column")

	// TEST6: sort by row.
	matrix = [][]float64{
		{1, 2, 3},
		{2, 3, 1},
		{3, 1, 2},
	}
	names = []string{"row1", "row2", "row3"}
	sortOrder = []string{"row2", "row3", "row1"}
	wantSorted = [][]float64{
		{2, 3, 1},
		{3, 1, 2},
		{1, 2, 3},
	}
	sorted, _ = Sort(matrix, names, sortOrder, "row")
	assert.Equal(t, wantSorted, sorted, "Matrix not sorted correctly by row")
}
