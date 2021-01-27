package leetgo

import (
	"fmt"
	"testing"
)

func Test_ladderLength(t *testing.T) {
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	b, e := "hit", "cog"
	fmt.Println(ladderLength(b, e, wordList))
}
