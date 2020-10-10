/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
package leetgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input    int
		expected int
	}{
		{12, 21},
		{-12, -21},
		{1, 1},
		{0, 0},
	}
	for _, item := range tests {
		assert.Equal(reverse(item.input), item.expected)
	}
}
