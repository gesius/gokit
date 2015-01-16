// Package unsafe implements wrapper for unsafe low level functions.
//
// For a higher level interface please use the github.com/pkg/term package.
package unsafe

import (
	"syscall"
	"unsafe"
)

type IOCTL struct{}

// Public interface.
func Run(fd uintptr, request int, status *int) error {
	return ioctl(fd, request, uintptr(unsafe.Pointer(status)))
}

//Native method call
func (s *IOCTL) ioctl(fd, request, argp uintptr) error {
	if _, _, e := syscall.Syscall6(syscall.SYS_IOCTL, fd, request, argp, 0, 0, 0); e != 0 {
		return e
	}
	return nil
}
