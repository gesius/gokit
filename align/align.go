package align

import "unsafe"

// Round the length  up to align it properly.
//see i.e for net link messages: https://go.googlesource.com/go/+/release-branch.go1.4/src/syscall/netlink_linux.go
func AlignOf(len int, alignto int) int {
	return (len + alignto - 1) & ^(alignto - 1)
}

func readUnaligned32(p unsafe.Pointer) uint32 {
	return *(*uint32)(p)
}
func readUnaligned64(p unsafe.Pointer) uint64 {
	return *(*uint64)(p)
}

func DIV_ROUND_UP(X int, Y int) int { return ((X) + ((Y) - 1)) / (Y) }
func ROUND_UP(X int, Y int) int     { return (DIV_ROUND_UP(X, Y)) * (Y) }
func PAD_SIZE(X int, Y int) int     { return (ROUND_UP(X, Y)) - (X) }
