/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */
package leetgo

import (
	"fmt"
	"testing"
)

func Test_searchRange(t *testing.T) {
	x := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(x, 8))
}
