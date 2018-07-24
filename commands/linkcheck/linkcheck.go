package linkcheck

import (
	"fmt"

	"github.com/tkiraly/lorawan/commands"
)

const Reqlen = 1

type LinkCheckReq interface {
	commands.Fopter
}

//LinkCheckReq command, an end-device may validate its connectivity with the network.
type linkCheckReq []byte

func (c linkCheckReq) ByteArray() []byte {
	return c
}

func (c linkCheckReq) Len() uint8 {
	return Reqlen
}

func (c linkCheckReq) String() string {
	return fmt.Sprintf("%s!", "LinkCheckReq")
}

func ParseReq(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return linkCheckReq(b)
}

func NewReq() LinkCheckReq {
	return linkCheckReq([]byte{commands.LinkCheckReqCommand})
}

const Anslen = 3

type LinkCheckAns interface {
	commands.Fopter
	Margin() uint8
	GwCnt() uint8
}

//LinkCheckReq command, an end-device may validate its connectivity with the network.
type linkCheckAns []byte

func (c linkCheckAns) ByteArray() []byte {
	return c
}

func (c linkCheckAns) Len() uint8 {
	return Anslen
}

func (c linkCheckAns) Margin() uint8 {
	return c[1]
}

func (c linkCheckAns) GwCnt() uint8 {
	return c[2]
}

func (c linkCheckAns) String() string {
	return fmt.Sprintf("%s! Margin: %d, GwCnt: %d;", "LinkCheckAns", c.Margin(), c.GwCnt())
}

func ParseAns(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return linkCheckAns(b)
}

func NewAns(margin, gwcnt uint8) LinkCheckAns {
	return linkCheckAns([]byte{commands.LinkCheckAnsCommand, margin, gwcnt})
}
