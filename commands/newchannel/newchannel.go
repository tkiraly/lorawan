package newchannel

import (
	"encoding/binary"
	"fmt"

	"github.com/tkiraly/lorawan/commands"
)

const Reqlen = 6

type NewChannelReq interface {
	commands.Fopter
	ChIndex() uint8
	Freq() uint32
	MaxDR() uint8
	MinDR() uint8
}

//NewChannelReq command, an end-device may validate its connectivity with the network.
type newChannelReq []byte

func (c newChannelReq) ByteArray() []byte {
	return c
}

func (c newChannelReq) Len() uint8 {
	return Reqlen
}

func (c newChannelReq) String() string {
	return fmt.Sprintf("%s! ChIndex: %d, Freq: %d, MaxDR: %d, MinDR: %d", "NewChannelReq",
		c.ChIndex(), c.Freq(), c.MaxDR(), c.MinDR())
}

func ParseReq(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return newChannelReq(b)
}

func NewReq(chindex uint8, freq uint32, maxdr, mindr uint8) NewChannelReq {
	temp := make([]byte, 4)
	binary.LittleEndian.PutUint32(temp, freq)
	return newChannelReq([]byte{commands.NewChannelReqCommand,
		chindex,
		temp[0],
		temp[1],
		temp[2],
		(maxdr << 4) | mindr,
	})
}

func (c newChannelReq) ChIndex() uint8 {
	return c[1]
}
func (c newChannelReq) Freq() uint32 {
	return binary.LittleEndian.Uint32([]byte{c[2], c[3], c[4], 0})
}
func (c newChannelReq) MaxDR() uint8 {
	return c[5] >> 4
}
func (c newChannelReq) MinDR() uint8 {
	return c[5] & 0x0F
}

const Anslen = 2

type NewChannelAns interface {
	commands.Fopter
	DataRateRangeOK() bool
	ChannelFrequencyOK() bool
}

//NewChannelReq command, an end-device may validate its connectivity with the network.
type newChannelAns []byte

func (c newChannelAns) ByteArray() []byte {
	return c
}

func (c newChannelAns) Len() uint8 {
	return Anslen
}

func (c newChannelAns) String() string {
	return fmt.Sprintf("%s! DataRateRangeOK: %t, ChannelFrequencyOK: %t", "NewChannelAns", c.DataRateRangeOK(), c.ChannelFrequencyOK())
}

func ParseAns(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return newChannelAns(b)
}

func NewAns(dataraterangeok, channelfrequencyok bool) NewChannelAns {
	v := byte(0)
	if dataraterangeok {
		v |= 0x02
	}
	if channelfrequencyok {
		v |= 0x01
	}
	return newChannelAns([]byte{commands.NewChannelAnsCommand,
		v,
	})
}

func (c newChannelAns) DataRateRangeOK() bool {
	return (c[1] & 0x02) == 0x02
}
func (c newChannelAns) ChannelFrequencyOK() bool {
	return (c[1] & 0x01) == 0x01
}
