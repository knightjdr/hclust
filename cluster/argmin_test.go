package cluster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgMin(t *testing.T) {
	// TEST1: find index of lowest data value not equal to index 3.
	data := []float64{0.5, 2, 3.3, 1.6}
	assert.Equal(t, 3, ArgMinGeneric(data, 1), "Not finding minimum value for anchor index for generic argmin")

	// TEST2: find index of lowest data value not equal to index 3.
	data = []float64{0.5, 2, 1.34, 0.2, 0.5}
	assert.Equal(t, 0, ArgMinNN(data, 3, -1), "Not finding minimum value for anchor index for NN argmin")

	// TEST3: find index of preferred element when two elements match min.
	assert.Equal(t, 4, ArgMinNN(data, 3, 4), "Not finding preferred index for NN argmin")

	// TEST5: find index of lowest data value.
	dataSingle := map[string]float64{
		"0": 0.5,
		"1": 2,
		"2": 1.34,
		"3": 0.2,
		"4": 0.5,
	}
	assert.Equal(t, "3", ArgMinSingle(dataSingle), "Not finding minimum value for single argmin")
}
