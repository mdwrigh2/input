package input

import (
	"unsafe"
	"github.com/mdwrigh2/ioctl"
)

var (
	EVIOCGVERSION = ioctl.IOR('E', 0x01, unsafe.Sizeof(int32(0)))
	EVIOCGID      = ioctl.IOR('E', 0x02, unsafe.Sizeof(InputId{}))
	EVIOCGKEYCODE = ioctl.IOR('E', 0x04, unsafe.Sizeof([2]uint{}))
	EVIOCGNAME    = func(nameSize int) int { return ioctl.IOR('E', 0x06, uintptr(nameSize)) }
	EVIOCGABS     = func(axis uint16) int {
		return ioctl.IOR('E', int(0x40+axis), unsafe.Sizeof(InputAbsInfo{}))
	}
)
