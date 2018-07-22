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

type FHDRDown interface {
	FHDR
	FCtrlDown
}

type FCtrlDown interface {
	FCtrl
	FPending() bool
}

type fHDRDown []byte

func (f fHDRDown) DevAddr() []byte {
	return util.Bytereverse(f[0:4])
}
func (f fHDRDown) FCnt() uint16 {
	cnt := binary.LittleEndian.Uint16(f[5:7])
	return cnt
}
func (f fHDRDown) FOpts() []commands.Fopter {
	if f.FOptsLen() > 0 {
		return ParseFOptsDown(f[7 : 7+f.FOptsLen()])
	}
	return nil
}
func (f fHDRDown) FPending() bool {
	return f[4]&0x10 == 0x10
}
func (f fHDRDown) ADR() bool {
	return f[4]&0x80 == 0x80
}
func (f fHDRDown) ACK() bool {
	return f[4]&0x20 == 0x20
}
func (f fHDRDown) FOptsLen() uint8 {
	return f[4] & 0x0f
}

func NewDown(devaddr []byte, fcnt uint16, fopts []commands.Fopter, adr,
	ack, fpending bool, foptslen uint8) FHDRDown {
	f := make([]byte, 7)
	copy(f[0:4], util.Bytereverse(devaddr))
	if adr {
		f[4] |= 0x80
	}
	if fpending {
		f[4] |= 0x10
	}
	if ack {
		f[4] |= 0x20
	}
	f[4] |= foptslen
	binary.LittleEndian.PutUint16(f[5:7], fcnt)
	for _, fopt := range fopts {
		f = append(f, fopt.ByteArray()...)
	}
	return fHDRDown(f)
}

func (f fHDRDown) String() string {
	return fmt.Sprintf("DevAddr: %s; FCnt: %d; ADR: %t; FPending: %t; ACK: %t; FOptsLen: %d, FOpts: %s\n",
		strings.ToUpper(hex.EncodeToString(f.DevAddr())),
		f.FCnt(),
		f.ADR(),
		f.FPending(),
		f.ACK(),
		f.FOptsLen(),
		f.FOpts())
}

func ParseDown(bb []byte) FHDRDown {
	b := make([]byte, len(bb))
	copy(b, bb)
	return fHDRDown(b)
}

func (f fHDRDown) ByteArray() []byte {
	return f
}

//ParseFOptsDown parses the MAC commands for an downlink frame
//
// NOTE: logs a warning to the console when unknown CID encountered
func ParseFOptsDown(payload []byte) []commands.Fopter {
	r := make([]commands.Fopter, 0)
	for i := 0; i < len(payload); {
		switch payload[i] {
		case commands.LinkCheckAnsCommand:
			r = append(r, linkcheck.ParseAns(payload[i:i+linkcheck.Anslen]))
			i += linkcheck.Anslen
		case commands.LinkADRReqCommand:
			r = append(r, linkadr.ParseReq(payload[i:i+linkadr.Reqlen]))
			i += linkadr.Reqlen
		case commands.DutyCycleReqCommand:
			r = append(r, dutycycle.ParseReq(payload[i:i+dutycycle.Reqlen]))
			i += dutycycle.Reqlen
		case commands.RXParamSetupReqCommand:
			r = append(r, rxparamsetup.ParseReq(payload[i:i+rxparamsetup.Reqlen]))
			i += rxparamsetup.Reqlen
		case commands.DevStatusReqCommand:
			r = append(r, devstatus.ParseReq(payload[i:i+devstatus.Reqlen]))
			i += devstatus.Reqlen
		case commands.NewChannelReqCommand:
			r = append(r, newchannel.ParseReq(payload[i:i+newchannel.Reqlen]))
			i += newchannel.Reqlen
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
