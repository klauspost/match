//+build !noasm
//+build !appengine

package match

import (
	"github.com/klauspost/cpuid"
)

func init() {
	hasAssembler = cpuid.CPU.SSE4()
}

func find4SSE4(needle, haystack []byte, dst []uint16)
func find4SSE4s(needle, haystack string, dst []uint16)
func find8SSE4(needle, haystack []byte, dst []uint32)
func find8SSE4s(needle, haystack string, dst []uint16)
