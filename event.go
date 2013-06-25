package input

import (
	"fmt"
	"unsafe"
	"syscall"
)

type EventType uint16
type EventCode uint16

const (
	// Event Types
	EV_SYN EventType = 0x00
	EV_KEY           = 0x01
	EV_REL           = 0x02
	EV_ABS           = 0x03
	EV_MSC           = 0x04
	EV_SW            = 0x05
	EV_LED           = 0x11
	EV_SND           = 0x12
	EV_REP           = 0x14
	EV_FF            = 0x15
	EV_PWR           = 0x16
	EV_FF_STATUS     = 0x17
	EV_MAX           = 0x1f
	EV_CNT           = (EV_MAX + 1)

	// Synchronization events
	SYN_REPORT    = 0
	SYN_CONFIG    = 1
	SYN_MT_REPORT = 2
	SYN_DROPPED   = 3
)

var (

	INPUT_EVENT_SIZE = int(unsafe.Sizeof(InputEvent{}))

	EventTypeToName = map[EventType]string{
		0x00: "EV_SYN",
		0x01: "EV_KEY",
		0x02: "EV_REL",
		0x03: "EV_ABS",
		0x04: "EV_MSC",
		0x05: "EV_SW",
		0x11: "EV_LED",
		0x12: "EV_SND",
		0x14: "EV_REP",
		0x15: "EV_FF",
		0x16: "EV_PWR",
		0x17: "EV_FF_STATUS",
		0x1f: "EV_MAX",
		0x20: "EV_CNT",
	}

	SyncCodeToName = map[EventCode]string{
		0: "SYN_REPORT",
		1: "SYN_CONFIG",
		2: "SYN_MT_REPORT",
		3: "SYN_DROPPED",
	}

	EventTypeToTranslation = map[EventType]map[EventCode]string {
		EV_SYN: SyncCodeToName,
		EV_KEY: KeycodeToName,
		EV_REL: RelCodeToName,
		EV_ABS: AbsCodeToName,
		EV_MSC: MscCodeToName,
		EV_SW: SwitchCodeToName,
		EV_LED: LedCodeToName,
		EV_SND: SoundCodeToName,
		EV_REP: map[EventCode]string{},
		EV_FF: ForceFeedbackCodeToName,
		EV_PWR: map[EventCode]string{},
		EV_FF_STATUS: ForceFeedbackCodeToName,
	}

)

type InputEvent struct {
	Time  syscall.Timeval
	Type  EventType
	Code  EventCode
	Value uint32
}

func (e InputEvent) String() string {
	return fmt.Sprintf("[%d.%d] Type %02d, Code %02d, Value %02d",
		e.Time.Sec, e.Time.Usec, e.Type, e.Code, e.Value)
}

func (e InputEvent) TranslatedString() string {
	t, ok := EventTypeToName[e.Type]
	if !ok {
		t = string(e.Type)
	}
	code, ok := EventTypeToTranslation[e.Type][e.Code]
	if !ok {
		code = string(e.Code)
	}

	return fmt.Sprintf("[%d.%d] Type %s, Code %s, Value %02d",
		e.Time.Sec, e.Time.Usec, t, code, e.Value)
}
