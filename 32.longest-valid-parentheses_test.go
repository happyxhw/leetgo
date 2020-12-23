/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */
package leetgo

import (
	"fmt"
	"testing"
)

func Test_longestValidParentheses(t *testing.T) {
	s := "(()"
	fmt.Println(longestValidParentheses(s))
}
