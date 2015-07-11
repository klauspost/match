//+build !amd64 noasm appengine

package match

func init() {
	hasAssembler = false
}

func find4SSE4(needle, haystack, dst []byte) {
	panic("assembler not enabled")
}

func find8SSE4(needle, haystack []byte, dst []uint16) {
	panic("assembler not enabled")
}

func find4SSE4s(needle, haystack string, dst []uint16) {
	panic("assembler not enabled")
}

func find8SSE4s(needle, haystack string, dst []uint16) {
	panic("assembler not enabled")
}
