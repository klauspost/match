package match

var hasAssembler bool

// Match4 will return start indeces of all matches of a 4 byte needle
// in a haystack that is a multiple of 16 in length.
// Indeces are returned ordered from index 0 and upwards.
func Match4(needle, haystack []byte, indices []int) []int {
	if len(needle) != 4 {
		panic("length not 4")
	}
	if len(haystack)&15 != 0 {
		panic("haystack must be dividable by 16")
	}
	dst := make([]uint16, len(haystack)/16)
	if indices == nil {
		indices = make([]int, 0, 10)
	}
	find4(needle, haystack, dst)
	for i, v := range dst {
		j := 0
		for v != 0 {
			if v&1 == 1 {
				indices = append(indices, i*16+j)
			}
			v >>= 1
			j++
		}
	}
	return indices
}

// Match4String performs the same operation as Match4 on strings
func Match4String(needle, haystack string, indices []int) []int {
	if len(needle) != 4 {
		panic("length not 4")
	}
	if len(haystack)&15 != 0 {
		panic("haystack must be dividable by 16")
	}
	dst := make([]uint16, len(haystack)/16)
	if indices == nil {
		indices = make([]int, 0, 10)
	}
	find4string(needle, haystack, dst)
	for i, v := range dst {
		j := 0
		for v != 0 {
			if v&1 == 1 {
				indices = append(indices, i*16+j)
			}
			v >>= 1
			j++
		}
	}
	return indices
}

func find4(needle, haystack []byte, dst []uint16) {
	if hasAssembler {
		find4SSE4(needle, haystack, dst)
		return
	}
	find4Go(needle, haystack, dst)
}

func find4string(needle, haystack string, dst []uint16) {
	if hasAssembler {
		find4SSE4s(needle, haystack, dst)
		return
	}
	find4Go([]byte(needle), []byte(haystack), dst)
}

// find4Go is the reference implementation that mimmics the SSE4
// implemenation.
func find4Go(needle, haystack []byte, dst []uint16) {
	end := uint(len(haystack) - 3)
	for i := uint(0); i < end; i++ {
		if needle[0] == haystack[i] {
			if needle[1] == haystack[i+1] && needle[2] == haystack[i+2] && needle[3] == haystack[i+3] {
				dst[i>>4] |= 1 << (i & 15)
			}
		}
	}
}

// Match8 will return start indeces of all matches of a 8 byte needle
// in a haystack that is a multiple of 16 in length.
// Indeces are returned ordered from index 0 and upwards.
func Match8(needle, haystack []byte, indices []int) []int {
	if len(needle) != 8 {
		panic("length not 8")
	}
	if len(haystack)&15 != 0 {
		panic("haystack must be dividable by 16")
	}
	dst := make([]uint32, len(haystack)/16)
	if indices == nil {
		indices = make([]int, 0, 10)
	}
	find8(needle, haystack, dst)
	for i, v := range dst {
		j := 0
		for v != 0 {
			if v&3 == 3 {
				indices = append(indices, i*16+j)
			}
			v >>= 2
			j++
		}
	}
	return indices
}

// Match4And8 will return start indeces of all matches of a 8 byte needle
// in a haystack that is a multiple of 16 in length.
// Matches for the first four bytes are returned in the first index, and 8
// byte matches are returned in the second. An index that is an 8 byte match will
// not be present in the 4-byte matches.
// Indeces are returned ordered from index 0 and upwards.
func Match8And4(needle, haystack []byte, indices8 []int, indices4 []int) ([]int, []int) {
	if len(needle) != 8 {
		panic("length not 8")
	}
	if len(haystack)&15 != 0 {
		panic("haystack must be dividable by 16")
	}
	dst := make([]uint32, len(haystack)/16)
	if indices8 == nil {
		indices8 = make([]int, 0, 10)
	} else {
		indices8 = indices8[:0]
	}
	if indices4 == nil {
		indices4 = make([]int, 0, 10)
	} else {
		indices4 = indices4[:0]
	}
	find8(needle, haystack, dst)
	for i, v := range dst {
		j := 0
		for v != 0 {
			if v&3 == 3 {
				indices8 = append(indices8, i*16+j)
			} else if v&1 == 1 {
				indices4 = append(indices4, i*16+j)
			}
			v >>= 2
			j++
		}
	}
	return indices8, indices4
}

func find8(needle, haystack []byte, dst []uint32) {
	if hasAssembler {
		find8SSE4(needle, haystack, dst)
		return
	}
	find8Go(needle, haystack, dst)
}

// find8Go is the reference implementation that mimmics the SSE4
// implemenation.
func find8Go(needle, haystack []byte, dst []uint32) {
	end := uint(len(haystack) - 7)
	for i := uint(0); i < end; i++ {
		if needle[0] == haystack[i] && needle[1] == haystack[i+1] && needle[2] == haystack[i+2] && needle[3] == haystack[i+3] {
			dst[i>>4] |= 1 << ((i & 15) << 1)
		}
		if needle[4] == haystack[i+4] && needle[5] == haystack[i+5] && needle[6] == haystack[i+6] && needle[7] == haystack[i+7] {
			dst[i>>4] |= 2 << ((i & 15) << 1)
		}
	}
}
