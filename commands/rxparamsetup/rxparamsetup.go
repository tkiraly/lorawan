package rxparamsetup

import (
	"encoding/binary"
	"fmt"

	"github.com/tkiraly/lorawan/commands"
)

const Reqlen = 5

type RXParamSetupReq interface {
	commands.Fopter
	RX1DRoffset() uint8
	RX2DataRate() uint8
	Frequency() uint32
}

//RXParamSetupReq command, an end-device may validate its connectivity with the network.
type rXParamSetupReq []byte

func (c rXParamSetupReq) ByteArray() []byte {
	return c
}

func (c rXParamSetupReq) Len() uint8 {
	return Reqlen
}

func (c rXParamSetupReq) String() string {
	return fmt.Sprintf("%s! RX1DRoffset: %d, RX2DataRate: %d, Frequency: %d", "RXParamSetupReq",
		c.RX1DRoffset(),
		c.RX2DataRate(),
		c.Frequency(),
	)
}

func (c rXParamSetupReq) RX1DRoffset() uint8 {
	return (c[1] >> 4) & 0x07
}
func (c rXParamSetupReq) RX2DataRate() uint8 {
	return c[1] & 0x0F
}
func (c rXParamSetupReq) Frequency() uint32 {
	return binary.LittleEndian.Uint32([]byte{c[2], c[3], c[4], 0})
}

func ParseReq(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return rXParamSetupReq(b)
}

func NewReq(rx1droffset, rx2datarate uint8, frequency uint32) RXParamSetupReq {
	temp := make([]byte, 4)
	binary.LittleEndian.PutUint32(temp, frequency)
	return rXParamSetupReq([]byte{commands.RXParamSetupReqCommand,
		(rx1droffset << 4) & rx2datarate,
		temp[0],
		temp[1],
		temp[2],
	})
}

const Anslen = 2

type RXParamSetupAns interface {
	commands.Fopter
	RX1DRoffsetACK() bool
	RX2DatarateACK() bool
	ChannelACK() bool
}

//RXParamSetupReq command, an end-device may validate its connectivity with the network.
type rXParamSetupAns []byte

func (c rXParamSetupAns) ByteArray() []byte {
	return c
}

func (c rXParamSetupAns) Len() uint8 {
	return Anslen
}

func (c rXParamSetupAns) String() string {
	return fmt.Sprintf("%s! RX1DRoffsetACK: %t, RX2DatarateACK: %t, ChannelACK: %t", "RXParamSetupAns",
		c.RX1DRoffsetACK(), c.RX2DatarateACK(), c.ChannelACK())
}

func ParseAns(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return rXParamSetupAns(b)
}

func NewAns(rx1droffsetack, rx2datarateack, channelack bool) RXParamSetupAns {
	temp := byte(0)
	if rx1droffsetack {
		temp |= 0x04
	}
	if rx2datarateack {
		temp |= 0x02
	}
	if channelack {
		temp |= 0x01
	}
	return rXParamSetupAns([]byte{commands.RXParamSetupAnsCommand, temp})
}

func (c rXParamSetupAns) RX1DRoffsetACK() bool {
	return (c[1] & 0x04) == 0x04
}
func (c rXParamSetupAns) RX2DatarateACK() bool {
	return (c[1] & 0x02) == 0x02
}
func (c rXParamSetupAns) ChannelACK() bool {
	return (c[1] & 0x01) == 0x01
}
