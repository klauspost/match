package match

import (
	"fmt"
	"testing"
)

var haystack = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

func TestMatch4(t *testing.T) {
	found := Match4(haystack[8:12], haystack, nil)
	expect := "[]int{8, 24}"
	got := fmt.Sprintf("%#v", found)
	if expect != got {
		t.Fatal("Expected", expect, "but got", got)
	}
}
