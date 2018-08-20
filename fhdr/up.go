package fhdr

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/tkiraly/lorawan/commands"
	"github.com/tkiraly/lorawan/commands/devstatus"
	"github.com/tkiraly/lorawan/commands/dutycycle"
	"github.com/tkiraly/lorawan/commands/linkadr"
	"github.com/tkiraly/lorawan/commands/linkcheck"
	"github.com/tkiraly/lorawan/commands/newchannel"
	"github.com/tkiraly/lorawan/commands/rxparamsetup"
	"github.com/tkiraly/lorawan/commands/rxtimingsetup"
	"github.com/tkiraly/lorawan/util"
)

type FHDRUp interface {
	FHDR
	FCtrlUp
}

type FCtrlUp interface {
	FCtrl
	ADRACKReq() bool
}

type fHDRUp []byte

func (f fHDRUp) DevAddr() []byte {
	return util.Bytereverse(f[0:4])
}
func (f fHDRUp) FCnt() uint16 {
	cnt := binary.LittleEndian.Uint16(f[5:7])
	return cnt
}
func (f fHDRUp) FOpts() []commands.Fopter {
	if f.FOptsLen() > 0 {
		return ParseFOptsUp(f[7 : 7+f.FOptsLen()])
	}
	return nil
}
func (f fHDRUp) ADRACKReq() bool {
	return f[4]&0x40 == 0x40
}
func (f fHDRUp) ADR() bool {
	return f[4]&0x80 == 0x80
}
func (f fHDRUp) ACK() bool {
	return f[4]&0x20 == 0x20
}
func (f fHDRUp) FOptsLen() uint8 {
	return f[4] & 0x0f
}

func NewUp(devaddr []byte, fcnt uint16, fopts []commands.Fopter, adr, ack, adrackreq bool, foptslen uint8) FHDRUp {
	f := make([]byte, 7)
	copy(f[0:4], util.Bytereverse(devaddr))
	if adr {
		f[4] |= 0x80
	}
	if adrackreq {
		f[4] |= 0x40
	}
	if ack {
		f[4] |= 0x20
	}
	f[4] |= foptslen
	binary.LittleEndian.PutUint16(f[5:7], fcnt)
	for _, fopt := range fopts {
		f = append(f, fopt.ByteArray()...)
	}
	return fHDRUp(f)
}

func (f fHDRUp) String() string {
	return fmt.Sprintf("DevAddr: %s; FCnt: %d; ADR: %t; ADRACKReq: %t; ACK: %t; FOptsLen: %d, FOpts: %s;",
		strings.ToUpper(hex.EncodeToString(f.DevAddr())),
		f.FCnt(),
		f.ADR(),
		f.ADRACKReq(),
		f.ACK(),
		f.FOptsLen(),
		f.FOpts())
}

func ParseUp(bb []byte) FHDRUp {
	b := make([]byte, len(bb))
	copy(b, bb)
	return fHDRUp(b)
}

func (f fHDRUp) ByteArray() []byte {
	return f
}

//ParseFOptsUp parses the MAC commands for an uplink frame
//
// NOTE: logs a warning to the console when unknown CID encountered
func ParseFOptsUp(payload []byte) []commands.Fopter {
	r := make([]commands.Fopter, 0)
	for i := 0; i < len(payload); {
		switch payload[i] {
		case commands.LinkCheckReqCommand:
			r = append(r, linkcheck.ParseReq(payload[i:i+linkcheck.Reqlen]))
			i += linkcheck.Reqlen
		case commands.LinkADRAnsCommand:
			r = append(r, linkadr.ParseAns(payload[i:i+linkadr.Anslen]))
			i += linkadr.Anslen
		case commands.DutyCycleAnsCommand:
			r = append(r, dutycycle.ParseAns(payload[i:i+dutycycle.Anslen]))
			i += dutycycle.Anslen
		case commands.RXParamSetupAnsCommand:
			r = append(r, rxparamsetup.ParseAns(payload[i:i+rxparamsetup.Anslen]))
			i += rxparamsetup.Anslen
		case commands.DevStatusAnsCommand:
			r = append(r, devstatus.ParseAns(payload[i:i+devstatus.Anslen]))
			i += devstatus.Anslen
		case commands.NewChannelAnsCommand:
			r = append(r, newchannel.ParseAns(payload[i:i+newchannel.Anslen]))
			i += newchannel.Anslen
		case commands.RxTimingSetupAnsCommand:
			r = append(r, rxtimingsetup.ParseAns(payload[i:i+rxtimingsetup.Anslen]))
			i += rxtimingsetup.Anslen
		default:
			log.Printf("WARNING: unknown MAC command id: %d, stopped parsing MAC commands", payload[i])
			return r
		}

	}
	return r
}
