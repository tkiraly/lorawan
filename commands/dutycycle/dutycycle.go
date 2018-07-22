package dutycycle

import (
	"fmt"

	"github.com/tkiraly/lorawan/commands"
)

const Reqlen = 2

type DutyCycleReq interface {
	commands.Fopter
	MaxDCycle() uint8
}

//DutyCycleReq command, an end-device may validate its connectivity with the network.
type dutyCycleReq []byte

func (c dutyCycleReq) ByteArray() []byte {
	return c
}

func (c dutyCycleReq) Len() uint8 {
	return Reqlen
}

func (c dutyCycleReq) String() string {
	return fmt.Sprintf("%s!", "DutyCycleReq")
}

func ParseReq(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return dutyCycleReq(b)
}

func NewReq(maxdcycle uint8) DutyCycleReq {
	return dutyCycleReq([]byte{commands.DutyCycleReqCommand,
		maxdcycle,
	})
}

func (c dutyCycleReq) MaxDCycle() uint8 {
	return c[1]
}

const Anslen = 1

type DutyCycleAns interface {
	commands.Fopter
}

//DutyCycleReq command, an end-device may validate its connectivity with the network.
type dutyCycleAns []byte

func (c dutyCycleAns) ByteArray() []byte {
	return c
}

func (c dutyCycleAns) Len() uint8 {
	return Reqlen
}

func (c dutyCycleAns) String() string {
	return fmt.Sprintf("%s!", "DutyCycleAns")
}

func ParseAns(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return dutyCycleAns(b)
}

func NewAns() DutyCycleAns {
	return dutyCycleAns([]byte{commands.DutyCycleAnsCommand})
}
