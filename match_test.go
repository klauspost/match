package match

import (
	"fmt"
	"reflect"
	"testing"

	fuzz "github.com/google/gofuzz"
)

func ExampleMatch4() {
	var haystack = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	}
	found := Match4(haystack[8:12], haystack, nil)
	fmt.Printf("%#v", found)
	//Output: []int{8, 24, 40, 56}
}

func TestMatch4End(t *testing.T) {
	var haystack = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	}
	found := Match4(haystack[60:64], haystack, nil)
	expect := "[]int{12, 28, 44, 60}"
	got := fmt.Sprintf("%#v", found)
	if expect != got {
		t.Fatal("Expected", expect, "but got", got)
	}
}

func ExampleMatch8() {
	var haystack = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
	}
	found := Match8(haystack[8:16], haystack, nil)
	fmt.Printf("%#v", found)
	// Output: []int{8, 24, 40, 56}
}

func ExampleMatch8And4() {
	var haystack = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115,
	}
	// Match 12 and 8 bytes
	f8, f4 := Match8And4(haystack[12:20], haystack, nil, nil)

	fmt.Printf("Length 8 match: %#v\n", f8)
	fmt.Printf("Length 4 match: %#v", f4)
	// Output: Length 8 match: []int{12}
	// Length 4 match: []int{28}
}

func ExampleMatchLen() {
	var data = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115,
	}
	// Get match length
	length := MatchLen(data[10:], data[26:], 16)

	fmt.Printf("Number of matching bytes: %d, first mismatch %d/%d\n", length, data[10+length], data[26+length])
	// Output: Number of matching bytes: 6, first mismatch 0/100
}

func TestMatch8And4s16(t *testing.T) {
	size := 16
	ta := make([]byte, size)
	found8, found4 := Match8And4(ta[0:8], ta, nil, nil)
	expect4 := []int{9, 10, 11, 12}
	expect8 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(found4, expect4) {
		t.Errorf("4matches: got \n%#v, expected:\n%#v", found4, expect4)
	}
	if !reflect.DeepEqual(found8, expect8) {
		t.Errorf("8matches: got \n%#v, expected:\n%#v", found8, expect8)
	}
	UseSse41 = false
	found8, found4 = Match8And4(ta[0:8], ta, nil, nil)
	if !reflect.DeepEqual(found4, expect4) {
		t.Errorf("4matches: got \n%#v, expected:\n%#v", found4, expect4)
	}
	if !reflect.DeepEqual(found8, expect8) {
		t.Errorf("8matches: got \n%#v, expected:\n%#v", found8, expect8)
	}
	UseSse41 = true
}

func TestMatch8And4s32(t *testing.T) {
	size := 32
	ta := make([]byte, size)
	found8, found4 := Match8And4(ta[0:8], ta, nil, nil)
	expect4 := []int{25, 26, 27, 28}
	expect8 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	if !reflect.DeepEqual(found4, expect4) {
		t.Errorf("4matches: got \n%#v, expected:\n%#v", found4, expect4)
	}
	if !reflect.DeepEqual(found8, expect8) {
		t.Errorf("8matches: got \n%#v, expected:\n%#v", found8, expect8)
	}
	UseSse41 = false
	found8, found4 = Match8And4(ta[0:8], ta, nil, nil)
	if !reflect.DeepEqual(found4, expect4) {
		t.Errorf("4matches: got \n%#v, expected:\n%#v", found4, expect4)
	}
	if !reflect.DeepEqual(found8, expect8) {
		t.Errorf("8matches: got \n%#v, expected:\n%#v", found8, expect8)
	}
	UseSse41 = true
}

func TestMatchLen(t *testing.T) {
	var data []byte

	for size := 0; size < 1000; size++ {
		f := fuzz.New()
		f.NumElements(size, size*2)
		f.NilChance(0.0)
		f.Fuzz(&data)

		length := MatchLen(data, data, size)
		if length != size {
			t.Fatalf("unexpected match length, (got) %d != %d (expected)", length, size)
		}
		if size == 0 {
			continue
		}
		// Change a value, and test if it is picked up
		var m int
		f.Fuzz(&m)
		if m < 0 {
			m *= -1
		}
		m %= len(data)
		var b = make([]byte, len(data))
		copy(b, data)
		b[m] = b[m] ^ 255
		length = MatchLen(data, b, size)
		if m < size {
			if length != m {
				t.Fatalf("unexpected match length, (got) %d != %d (expected)", length, m)
			}
		} else {
			if length != size {
				t.Fatalf("unexpected match length, (got) %d != %d (expected)", length, size)
			}
		}
	}
}

func BenchmarkMatch4String(b *testing.B) {
	size := 1024
	found := make([]int, 0, 10)
	ta := make([]byte, size)
	f := fuzz.New()
	f.NumElements(size, size)
	f.NilChance(0.0)
	f.Fuzz(&ta)
	txt := string(ta)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found = Match4String(txt[800:804], txt, found)
	}
}

// Shows the overhead of converting to bytes.
func BenchmarkMatch4Convert(b *testing.B) {
	size := 1024
	found := make([]int, 0, 10)
	ta := make([]byte, size)
	f := fuzz.New()
	f.NumElements(size, size)
	f.NilChance(0.0)
	f.Fuzz(&ta)
	txt := string(ta)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found = Match4([]byte(txt[800:804]), []byte(txt), found)
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

func BenchmarkMatchLen256(b *testing.B) {
	size := 256
	ta := make([]byte, size)
	f := fuzz.New()
	f.NumElements(size, size)
	f.NilChance(0.0)
	f.Fuzz(&ta)
	b.SetBytes(int64(size))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MatchLen(ta, ta, size)
	}
}
