package leetgo

import (
	"fmt"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	x := []int{}
	fmt.Println("start")
	for i := 10000; i > -10000; i-- {
		x = append(x, i)
	}
	res := maxSlidingWindow(x, 10000)
	for _, item := range res {
		fmt.Println(item)
	}
}
