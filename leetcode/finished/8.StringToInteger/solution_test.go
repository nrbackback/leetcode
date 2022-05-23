package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(-42, myAtoi("   -42"))
	assert.Equal(4193, myAtoi("4193 with words"))
	assert.Equal(0, myAtoi("words and 987"))
	assert.Equal(-2147483648, myAtoi("-91283472332"))
	// 提交后的单元测试

	assert.Equal(1, myAtoi("+1"))
	assert.Equal(-2147483648, myAtoi("-2147483648"))

	// 需要考虑越界情况
	assert.Equal(2147483647, myAtoi("20000000000000000000"))
	assert.Equal(-2147483648, myAtoi("-9223372036854775809"))
}
