package leetgo

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {
	arg := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate_48(arg)
	fmt.Println(arg)
}
