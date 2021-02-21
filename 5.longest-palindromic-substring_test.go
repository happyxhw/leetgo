/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */
package leetgo

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	x := "babad"
	fmt.Println(longestPalindrome(x))
}
