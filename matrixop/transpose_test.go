package matrixop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranpose(t *testing.T) {
	matrix := [][]float64{
		{5, 2, 14.3, 2.1},
		{23, 17.8, 0, 0.4},
		{10, 0, 7, 15.9},
	}

	// TEST1: tranpose matrix.
	want := [][]float64{
		{5, 23, 10},
		{2, 17.8, 0},
		{14.3, 0, 7},
		{2.1, 0.4, 15.9},
	}
	assert.Equal(t, want, Transpose(matrix), "Matrix not tranposed correctly")
}
