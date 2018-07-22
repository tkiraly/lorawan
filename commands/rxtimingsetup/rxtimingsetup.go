package rxtimingsetup

import (
	"fmt"

	"github.com/tkiraly/lorawan/commands"
)

const Reqlen = 2

type RxTimingSetupReq interface {
	commands.Fopter
	Del() uint8
}

//RxTimingSetupReq command, an end-device may validate its connectivity with the network.
type rxTimingSetupReq []byte

func (c rxTimingSetupReq) ByteArray() []byte {
	return c
}

func (c rxTimingSetupReq) Len() uint8 {
	return Reqlen
}

func (c rxTimingSetupReq) String() string {
	return fmt.Sprintf("%s!", "RxTimingSetupReq")
}

func ParseReq(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return rxTimingSetupReq(b)
}

func NewReq(del uint8) RxTimingSetupReq {
	return rxTimingSetupReq([]byte{commands.RxTimingSetupReqCommand,
		del,
	})
}
func (c rxTimingSetupReq) Del() uint8 {
	return c[1]
}

const Anslen = 1

type RxTimingSetupAns interface {
	commands.Fopter
}

//RxTimingSetupReq command, an end-device may validate its connectivity with the network.
type rxTimingSetupAns []byte

func (c rxTimingSetupAns) ByteArray() []byte {
	return c
}

func (c rxTimingSetupAns) Len() uint8 {
	return Reqlen
}

func (c rxTimingSetupAns) String() string {
	return fmt.Sprintf("%s!", "RxTimingSetupAns")
}

func ParseAns(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return rxTimingSetupAns(b)
}

func NewAns() RxTimingSetupAns {
	return rxTimingSetupAns([]byte{commands.RxTimingSetupAnsCommand})
}
