package match

import (
	"fmt"
	"testing"
)

var haystack = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

func TestMatch4(t *testing.T) {
	found := Match4(haystack[8:12], haystack, nil)
	expect := "[]int{8, 24, 40, 56}"
	got := fmt.Sprintf("%#v", found)
	if expect != got {
		t.Fatal("Expected", expect, "but got", got)
	}
}

func TestMatch4End(t *testing.T) {
	found := Match4(haystack[60:64], haystack, nil)
	expect := "[]int{12, 28, 44, 60}"
	got := fmt.Sprintf("%#v", found)
	if expect != got {
		t.Fatal("Expected", expect, "but got", got)
	}
}

func TestMatch8(t *testing.T) {
	found := Match8(haystack[8:16], haystack, nil)
	expect := "[]int{8, 24, 40, 56}"
	got := fmt.Sprintf("%#v", found)
	if expect != got {
		t.Fatal("Expected", expect, "but got", got)
	}
}
