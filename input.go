package input

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"unsafe"
	"github.com/mdwrigh2/ioctl"
)

type InputId struct {
	Bustype, Vendor, Product, Version uint16
}

type InputAbsInfo struct {
	Value, Min, Max, Fuzz, Flat, Resolution int32
}


func GetDeviceId(file *os.File) (*InputId, error) {
	inputId := new(InputId)
	if err := ioctl.Ioctl(file.Fd(), EVIOCGID, unsafe.Pointer(inputId)); err != 0 {
		return nil, err
	}
	return inputId, nil
}

func GetDeviceDriverVersion(file *os.File) (int32, error) {
	version := new(int32)
	if err := ioctl.Ioctl(file.Fd(), EVIOCGVERSION, unsafe.Pointer(version)); err != 0 {
		return 0, err
	}
	return *version, nil
}

func GetDeviceName(file *os.File) (string, error) {
	var name [80]byte
	if err := ioctl.Ioctl(file.Fd(), EVIOCGNAME(79), unsafe.Pointer(&name)); err != 0 {
		return "", err
	}
	runes := bytes.Runes(name[:])
	var i int
	for i = range runes {
		if !strconv.IsPrint(runes[i]) {
			break
		}
	}
	nameString := strings.TrimSpace(string(runes[:i]))
	return nameString, nil
}

func GetAxisInfo(file *os.File, axis uint16) (*InputAbsInfo, error) {
	info := new(InputAbsInfo)
	if err := ioctl.Ioctl(file.Fd(), EVIOCGABS(axis), unsafe.Pointer(info)); err != 0 {
		return nil, err
	}
	return info, nil
}
