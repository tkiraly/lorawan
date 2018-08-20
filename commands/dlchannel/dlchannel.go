package dlchannel

import (
	"encoding/binary"
	"fmt"

	"github.com/tkiraly/lorawan/commands"
)

const Reqlen = 5

type DlChannelReq interface {
	commands.Fopter
	ChIndex() uint8
	Freq() uint32
}

type dlChannelReq []byte

func (c dlChannelReq) ByteArray() []byte {
	return c
}

func (c dlChannelReq) Len() uint8 {
	return Reqlen
}

func (c dlChannelReq) String() string {
	return fmt.Sprintf("%s! ChIndex: %d, Freq: %d", "DlChannelReq",
		c.ChIndex(),
		c.Freq())
}

func ParseReq(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return dlChannelReq(b)
}

func NewReq(chindex uint8, freq uint32) DlChannelReq {
	temp := make([]byte, 4)
	binary.LittleEndian.PutUint32(temp, freq)
	return dlChannelReq([]byte{commands.DlChannelReqCommand,
		chindex,
		temp[0],
		temp[1],
		temp[2],
	})
}

func (c dlChannelReq) ChIndex() uint8 {
	return c[1]
}
func (c dlChannelReq) Freq() uint32 {
	return binary.LittleEndian.Uint32([]byte{c[2], c[3], c[4], 0})
}

const Anslen = 2

type DlChannelAns interface {
	commands.Fopter
	UplinkFrequencyExists() bool
	ChannelFrequencyOk() bool
}

type dlChannelAns []byte

func (c dlChannelAns) ByteArray() []byte {
	return c
}

func (c dlChannelAns) Len() uint8 {
	return Anslen
}

func (c dlChannelAns) String() string {
	return fmt.Sprintf("%s! UplinkFrequencyExists: %t, ChannelFrequencyOk: %t",
		"DlChannelAns",
		c.UplinkFrequencyExists(),
		c.ChannelFrequencyOk())
}

func ParseAns(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return dlChannelAns(b)
}

func NewAns(uplinkfrequencyexists, channelfrequencyok bool) DlChannelAns {
	var temp = byte(0)
	if uplinkfrequencyexists {
		temp |= 0x02
	}
	if channelfrequencyok {
		temp |= 0x01
	}
	return dlChannelAns([]byte{commands.DlChannelAnsCommand,
		temp,
	})
}

func (c dlChannelAns) ChannelFrequencyOk() bool {
	return (c[1]>>1)&0x01 == 0x01
}
func (c dlChannelAns) UplinkFrequencyExists() bool {
	return c[1]&0x01 == 0x01
}
