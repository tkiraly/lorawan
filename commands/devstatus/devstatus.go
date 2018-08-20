package devstatus

import (
	"fmt"

	"github.com/tkiraly/lorawan/commands"
)

// Reqlen is the length of the request
const Reqlen = 1

// DevStatusReq requests for the status of the end-device
// With the DevStatusReq command a network server
// may request status information from an end-device.
type DevStatusReq interface {
	commands.Fopter
}

type devStatusReq []byte

func (c devStatusReq) ByteArray() []byte {
	return c
}

func (c devStatusReq) Len() uint8 {
	return Reqlen
}

func (c devStatusReq) String() string {
	return fmt.Sprintf("%s!", "DevStatusReq")
}

// ParseReq parses the byte array representation of the request
func ParseReq(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return devStatusReq(b)
}

// NewReq creates a new request
func NewReq() DevStatusReq {
	return devStatusReq([]byte{commands.DevStatusReqCommand})
}

// Anslen is the length of the answer
const Anslen = 3

// DevStatusAns contains the status of the end-device
// If a DevStatusReq is received by an end-device,
// it responds with a DevStatusAns command.
// +======================+=========+========+
// |     Size (bytes)     |    1    |   1    |
// +======================+=========+========+
// | DevStatusAns Payload | Battery | Margin |
// +----------------------+---------+--------+
type DevStatusAns interface {
	commands.Fopter
	// +=========+================================================================+
	// | Battery |                          Description                           |
	// +=========+================================================================+
	// |       0 | The end-device is connected to an external power source.       |
	// +---------+----------------------------------------------------------------+
	// | 1...254 | The battery level, 1 being at minimum and 254 being at maximum |
	// +---------+----------------------------------------------------------------+
	// |     255 | The end-device was not able to measure the battery level.      |
	// +---------+----------------------------------------------------------------+
	Battery() uint8
	// The margin (Margin) is the demodulation signal-to-noise ratio in dB
	// rounded to the nearest integer value for the last successfully received
	// DevStatusReq command. It is a signed integer of 6 bits with a minimum
	// value of -32 and a maximum value of 31.
	// +========+=====+========+
	// |  Bits  | 7:6 |  5:0   |
	// +========+=====+========+
	// | Status | RFU | Margin |
	// +--------+-----+--------+
	Margin() uint8
}

type devStatusAns []byte

func (c devStatusAns) ByteArray() []byte {
	return c
}

func (c devStatusAns) Len() uint8 {
	return Anslen
}

func (c devStatusAns) Battery() uint8 {
	return c[1]
}

func (c devStatusAns) Margin() uint8 {
	return c[2]
}

func (c devStatusAns) String() string {
	return fmt.Sprintf("%s! Battery: %d, Margin: %d", "DevStatusAns", c.Battery(), c.Margin())
}

// ParseAns parses the byte array representation of the answer
func ParseAns(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return devStatusAns(b)
}

// NewAns creates a new answer
func NewAns(battery, margin uint8) DevStatusAns {
	return devStatusAns([]byte{commands.DevStatusAnsCommand, battery, margin})
}
