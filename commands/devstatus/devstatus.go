package devstatus

import (
	"fmt"

	"github.com/tkiraly/lorawan/commands"
)

const Reqlen = 1

type DevStatusReq interface {
	commands.Fopter
}

//DevStatusReq command, an end-device may validate its connectivity with the network.
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

func ParseReq(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return devStatusReq(b)
}

func NewReq() DevStatusReq {
	return devStatusReq([]byte{commands.DevStatusReqCommand})
}

const Anslen = 3

type DevStatusAns interface {
	commands.Fopter
	Battery() uint8
	Margin() uint8
}

//DevStatusReq command, an end-device may validate its connectivity with the network.
type devStatusAns []byte

func (c devStatusAns) ByteArray() []byte {
	return c
}

func (c devStatusAns) Len() uint8 {
	return Reqlen
}

func (c devStatusAns) Battery() uint8 {
	return c[1]
}

func (c devStatusAns) Margin() uint8 {
	return c[2]
}

func (c devStatusAns) String() string {
	return fmt.Sprintf("%s! Battery: %d, Margin: %d;", "DevStatusAns", c.Battery(), c.Margin())
}

func ParseAns(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return devStatusAns(b)
}

func NewAns(battery, margin uint8) DevStatusAns {
	return devStatusAns([]byte{commands.DevStatusAnsCommand, battery, margin})
}
