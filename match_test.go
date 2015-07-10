package match

import (
	"fmt"
	"testing"

	fuzz "github.com/google/gofuzz"
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

func BenchmarkMatch8(b *testing.B) {
	size := 32768
	ta := make([]byte, size)
	found := make([]int, 0, 10)
	f := fuzz.New()
	f.NumElements(size, size)
	f.NilChance(0.0)
	f.Fuzz(&ta)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found = Match8(ta[800:808], ta, found)
	}
}

func BenchmarkMatch8And4(b *testing.B) {
	size := 32768
	ta := make([]byte, size)
	found4 := make([]int, 0, 10)
	found8 := make([]int, 0, 10)
	f := fuzz.New()
	f.NumElements(size, size)
	f.NilChance(0.0)
	f.Fuzz(&ta)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found4, found8 = Match8And4(ta[800:808], ta, found4, found8)
	}
}
