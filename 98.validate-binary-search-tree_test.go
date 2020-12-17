package leetgo

import (
	"fmt"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	x := TreeNode{Val: 1}
	y := TreeNode{Val: 1}
	// z := TreeNode{Val: 3}
	x.Left = &y
	// x.Right = &z

	fmt.Println(isValidBST(&x))
}
