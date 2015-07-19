//+build !amd64 noasm appengine

package match

func init() {
	UseSse41 = false
	UseSse42 = false
}

func find4SSE4(needle, haystack []byte, dst []uint16) {
	panic("assembler not enabled")
}

func find8SSE4(needle, haystack []byte, dst []uint32) {
	panic("assembler not enabled")
}

func find4SSE4s(needle, haystack string, dst []uint16) {
	panic("assembler not enabled")
}

func find8SSE4s(needle, haystack string, dst []uint32) {
	panic("assembler not enabled")
}

func matchLenSSE4(a, b []byte, max int) int {
	panic("assembler not enabled")
}
