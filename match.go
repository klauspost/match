package match

import "fmt"

func Match4(needle, haystack []byte, indices []int) []int {
	dst := make([]byte, len(haystack)/8)
	if indices == nil {
		indices = make([]int, 0, 10)
	}
	if len(needle) != 4 {
		panic("length not 4")
	}
	if len(haystack)&15 != 0 {
		panic("haystack must be dividable by 16")
	}
	Find4(needle, haystack, dst)
	for i, v := range dst {
		j := 0
		for v != 0 {
			if v&1 == 1 {
				indices = append(indices, i*8+j)
			}
			v >>= 1
			j++
		}
	}
	return indices
}

func Find4(needle, haystack, dst []byte) {
	if true {
		Find4SSE4(needle, haystack, dst)
		return
	}
	end := uint(len(haystack) - 4)
	for i := uint(0); i < end; i++ {
		if needle[0] == haystack[i] {
			if needle[1] == haystack[i+1] && needle[2] == haystack[i+2] && needle[3] == haystack[i+3] {
				dst[i>>3] |= 1 << (i & 7)
			}
		}
	}
}

func Match8(needle, haystack []byte) []int {
	return nil
}
