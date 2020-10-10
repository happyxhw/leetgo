/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
package leetgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input    string
		expected int
	}{
		{"aba", 2},
		{"ab", 2},
		{"a", 1},
		{"abacade", 4},
	}
	for _, item := range tests {
		assert.Equal(lengthOfLongestSubstring(item.input), item.expected)
	}
}
