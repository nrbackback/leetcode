package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	assert := assert.New(t)

	s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	r := groupAnagrams(s)
	assert.Equal([][]string{{"bat"}, {"tan", "nat"}, {"eat", "tea", "ate"}}, r)
}
