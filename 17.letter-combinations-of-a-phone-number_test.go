/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
package leetgo

import (
	"fmt"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	x := "23"
	fmt.Println(letterCombinations(x))
	// ["ad","ae","af","bd","be","bf","cd","ce","cf"]
}
