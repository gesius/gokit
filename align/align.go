package align

import (
	. "unsafe"
)


func(int, int)int { return x + y }

func DIV_ROUND_UP(X int,Y int) int {return ((X) + ((Y) - 1)) / (Y)}
func ROUND_UP(X int,Y int) int {return (DIV_ROUND_UP(X, Y)) * (Y)}
func PAD_SIZE(X int,Y int) int {return (ROUND_UP(X, Y)) - (X)}

"
func readUnaligned32(p unsafe.Pointer) uint32 {
	return readUnaligned32(p)
	
	//return *(*uint32)(p)
}
func readUnaligned64(p unsafe.Pointer) uint64 {
	return *(*uint64)(p)
}