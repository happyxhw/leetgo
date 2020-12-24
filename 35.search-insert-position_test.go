/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */
package leetgo

import (
	"fmt"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	x := []int{1, 3}
	fmt.Println(searchInsert(x, 2))
}
