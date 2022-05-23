package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxMatrixSum(t *testing.T) {
	assert := assert.New(t)

	m := [][]int{{1, -1}, {-1, 1}}
	assert.Equal(int64(4), maxMatrixSum(m))

	m = [][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}
	assert.Equal(int64(16), maxMatrixSum(m))
}
