package linkadr

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/tkiraly/lorawan/commands"
)

const Reqlen = 5

type LinkADRReq interface {
	commands.Fopter
	DataRate() uint8
	TxPower() uint8
	ChMask() uint16
	Redundancy_ChMaskCntl() uint8
	Redundancy_NbTrans() uint8
}

//LinkADRReq command, an end-device may validate its connectivity with the network.
type linkADRReq []byte

func (c linkADRReq) ByteArray() []byte {
	return c
}

func (c linkADRReq) Len() uint8 {
	return Reqlen
}

func (c linkADRReq) String() string {
	return fmt.Sprintf("%s! Datarate: %d, TxPower: %d, Chmask: %s, ChmaskCtrl: %d, Nbtrans: %d",
		"LinkADRReq", c.DataRate(), c.TxPower(), strconv.FormatUint(uint64(c.ChMask()), 16), c.Redundancy_ChMaskCntl(), c.Redundancy_NbTrans())
}

func ParseReq(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return linkADRReq(b)
}

func NewReq(datarate, txpower uint8, chmask uint16, chmaskcntl, nbtrans uint8) LinkADRReq {
	return linkADRReq([]byte{commands.LinkADRReqCommand,
		(datarate << 4) | txpower,
		byte(chmask >> 8),
		byte(chmask),
		(chmaskcntl << 4) | nbtrans,
	})
}

func (c linkADRReq) DataRate() uint8 {
	return (c[1] >> 4)
}
func (c linkADRReq) TxPower() uint8 {
	return c[1] & 0x0F
}
func (c linkADRReq) ChMask() uint16 {
	return binary.BigEndian.Uint16(c[2:4])
}
func (c linkADRReq) Redundancy_ChMaskCntl() uint8 {
	return (c[4] >> 4) & 0x07
}
func (c linkADRReq) Redundancy_NbTrans() uint8 {
	return c[4] & 0x0F
}

const Anslen = 2

type LinkADRAns interface {
	commands.Fopter
	PowerACK() bool
	DatarateACK() bool
	ChannelmaskACK() bool
}

//LinkADRReq command, an end-device may validate its connectivity with the network.
type linkADRAns []byte

func (c linkADRAns) ByteArray() []byte {
	return c
}

func (c linkADRAns) Len() uint8 {
	return Anslen
}

func (c linkADRAns) PowerACK() bool {
	return (c[1] & 0x04) == 0x04
}
func (c linkADRAns) DatarateACK() bool {
	return (c[1] & 0x02) == 0x02
}
func (c linkADRAns) ChannelmaskACK() bool {
	return (c[1] & 0x01) == 0x01
}

func (c linkADRAns) String() string {
	return fmt.Sprintf("%s! PowerACK: %t, DatarateACK: %t, ChannelmaskACK: %t", "LinkADRAns",
		c.PowerACK(), c.DatarateACK(), c.ChannelmaskACK())
}

func ParseAns(bb []byte) commands.Fopter {
	b := make([]byte, len(bb))
	copy(b, bb)
	return linkADRAns(b)
}

func NewAns(powerack, datarateack, channelmaskack bool) LinkADRAns {
	temp := byte(0x00)
	if powerack {
		temp |= 0x04
	}
	if datarateack {
		temp |= 0x02
	}
	if channelmaskack {
		temp |= 0x01
	}

	return linkADRAns([]byte{commands.LinkADRAnsCommand, temp})
}
