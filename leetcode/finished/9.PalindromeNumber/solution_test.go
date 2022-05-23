package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(true, isPalindrome(121))
	assert.Equal(false, isPalindrome(-121))
	assert.Equal(false, isPalindrome(10))
	assert.Equal(false, isPalindrome(-101))
	// 提交运行失败后补充的测试用例
	assert.Equal(false, isPalindrome(1000021))
	assert.Equal(true, isPalindrome(11))
}
