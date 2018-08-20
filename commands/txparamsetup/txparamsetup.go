package txparamsetup

import (
	"fmt"

	"github.com/tkiraly/lorawan/commands"
)

const Reqlen = 2

type TxParamSetupReq interface {
	commands.Fopter
	DownlinkDwellTime() uint8
	UplinkDwellTime() uint8
	MaxEIRP() uint8
}

type txParamSetupReq []byte

func (c txParamSetupReq) ByteArray() []byte {
	return c
}

func (c txParamSetupReq) Len() uint8 {
	return Reqlen
}

func (c txParamSetupReq) String() string {
	return fmt.Sprintf("%s! DownlinkDwellTime: %d, UplinkDwellTime: %d, MaxEIRP: %d;",
		"TxParamSetupReq",
		c.DownlinkDwellTime(),
		c.UplinkDwellTime(),
		c.MaxEIRP())
}

func ParseReq(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return txParamSetupReq(b)
}

func NewReq(downlinkdwelltime, uplinkdwelltime, maxeirp uint8) TxParamSetupReq {
	return txParamSetupReq([]byte{commands.TxParamSetupReqCommand,
		(downlinkdwelltime << 5) |
			(uplinkdwelltime << 4) |
			maxeirp,
	})
}

func (c txParamSetupReq) DownlinkDwellTime() uint8 {
	return (c[1] >> 5) & 0x01
}
func (c txParamSetupReq) UplinkDwellTime() uint8 {
	return (c[1] >> 4) & 0x01

}
func (c txParamSetupReq) MaxEIRP() uint8 {
	return c[1] & 0x0f
}

const Anslen = 1

type TxParamSetupAns interface {
	commands.Fopter
}

type txParamSetupAns []byte

func (c txParamSetupAns) ByteArray() []byte {
	return c
}

func (c txParamSetupAns) Len() uint8 {
	return Anslen
}

func (c txParamSetupAns) String() string {
	return fmt.Sprintf("%s!", "TxParamSetupAns")
}

func ParseAns(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return txParamSetupAns(b)
}

func NewAns() TxParamSetupAns {
	return txParamSetupAns([]byte{commands.TxParamSetupAnsCommand})
}
