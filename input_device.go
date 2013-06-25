package input

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type InputDevice struct {
	Id       InputId
	Version  int32
	Name     string
	File     *os.File
	AxisInfo map[uint16]*InputAbsInfo
}

func NewInputDevice(file *os.File) (*InputDevice, error) {
	dev := new(InputDevice)
	dev.File = file
	if id, err := GetDeviceId(file); err != nil {
		return nil, fmt.Errorf("Failed to get device ID: %s", err.Error())
	} else {
		dev.Id = *id
	}
	if version, err := GetDeviceDriverVersion(file); err != nil {
		return nil, fmt.Errorf("Failed to get device driver version: %s", err.Error())
	} else {
		dev.Version = version
	}
	if name, err := GetDeviceName(file); err != nil {
		return nil, fmt.Errorf("Failed to get device name: %s", err.Error())
	} else {
		dev.Name = name
	}
	dev.AxisInfo = make(map[uint16]*InputAbsInfo)
	return dev, nil
}

func (dev *InputDevice) PopulateAxisInfo(axis uint16) (*InputAbsInfo, error) {
	if axis < 0 || axis > ABS_MAX {
		return nil, fmt.Errorf("Invalid axis id")
	}

	info, err := GetAxisInfo(dev.File, axis)
	if err != nil {
		return nil, err
	}
	dev.AxisInfo[axis] = info
	return info, nil
}

func (dev InputDevice) NextEvent() (*InputEvent, error) {
	eventBuffer := make([]byte, INPUT_EVENT_SIZE)
	if nbytes, err := dev.File.Read(eventBuffer); nbytes != INPUT_EVENT_SIZE || err != nil {
		fmt.Println(nbytes)
		var errString string
		if err != nil {
			errString = err.Error()
		} else {
			errString = fmt.Sprintf("Read incorrect number of bytes (%d != %d)", nbytes,
				INPUT_EVENT_SIZE)
		}
		return nil, fmt.Errorf("Failed to read event: %s", errString)
	}

	event := new(InputEvent)
	if err := binary.Read(bytes.NewBuffer(eventBuffer), binary.LittleEndian, event); err != nil {
		return nil, fmt.Errorf("Failed to read event: %s", err.Error())
	}

	return event, nil
}
