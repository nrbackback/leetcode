package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumAbsDifference(t *testing.T) {
	assert := assert.New(t)

	input := []int{4, 2, 1, 3}
	out := minimumAbsDifference(input)
	assert.Equal([][]int{{1, 2}, {2, 3}, {3, 4}}, out)

	input = []int{1, 3, 6, 10, 15}
	out = minimumAbsDifference(input)
	assert.Equal([][]int{{1, 3}}, out)

	input = []int{3, 8, -10, 23, 19, -4, -14, 27}
	out = minimumAbsDifference(input)
	assert.Equal([][]int{{-14, -10}, {19, 23}, {23, 27}}, out)

	// assert.False(true)
}
