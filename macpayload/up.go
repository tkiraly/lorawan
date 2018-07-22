package macpayload

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/tkiraly/lorawan/commands"
	"github.com/tkiraly/lorawan/fhdr"
	"github.com/tkiraly/lorawan/mhdr"
)

type DataUp interface {
	MType() mhdr.MType
	Major() mhdr.MajorVersion
	fhdr.FHDRUp
	FPort() *uint8
	FRMPayload() []byte
	MIC() []byte
}
type dataUp []byte

func (m dataUp) ByteArray() []byte {
	return m
}

func (m dataUp) MType() mhdr.MType {
	v := mhdr.Parse(m[0])
	return v.MType()
}

func (m dataUp) Major() mhdr.MajorVersion {
	v := mhdr.Parse(m[0])
	return v.Major()
}

func (d dataUp) FPort() *uint8 {
	f := fhdr.ParseDown(d[1:])
	i := 1 + 7 + int(f.FOptsLen()) + 4
	if len(d) > i {
		return &d[1+7+f.FOptsLen()]
	}
	return nil
}

func (d dataUp) FRMPayload() []byte {
	f := fhdr.ParseDown(d[1:])
	i := 1 + 7 + int(f.FOptsLen()) + 4
	if len(d) > i {
		return d[1+7+f.FOptsLen()+1 : len(d)-4]
	}
	return nil
}

func (d dataUp) MIC() []byte {
	return d[len(d)-4:]
}

func (du dataUp) String() string {
	var fport *byte
	if du.FPort() != nil {
		fport = du.FPort()
	}
	return fmt.Sprintf("DataDown: MHDR: %s; FHDRUp: %s; FPort: %d; FRMPayload: %s; MIC: %s\n",
		mhdr.Parse(du[0]),
		fhdr.ParseDown(du[1:]),
		*fport,
		strings.ToUpper(hex.EncodeToString(du.FRMPayload())),
		strings.ToUpper(hex.EncodeToString(du.MIC())))
}
func (du dataUp) DevAddr() []byte {
	d := fhdr.ParseUp(du[1:])
	return d.DevAddr()
}
func (du dataUp) FCnt() uint16 {
	d := fhdr.ParseUp(du[1:])
	return d.FCnt()
}
func (du dataUp) FOpts() []commands.Fopter {
	d := fhdr.ParseUp(du[1:])
	return d.FOpts()
}
func (du dataUp) ADRACKReq() bool {
	d := fhdr.ParseUp(du[1:])
	return d.ADRACKReq()
}
func (du dataUp) ADR() bool {
	d := fhdr.ParseUp(du[1:])
	return d.ADR()
}
func (du dataUp) ACK() bool {
	d := fhdr.ParseUp(du[1:])
	return d.ACK()
}
func (du dataUp) FOptsLen() uint8 {
	d := fhdr.ParseUp(du[1:])
	return d.FOptsLen()
}

func NewUp(mtype mhdr.MType, major mhdr.MajorVersion, devaddr []byte,
	fcnt uint16, fopts []commands.Fopter, adr,
	ack, adrackreq bool, foptslen uint8, fport *uint8,
	frmpayload, mic []byte) DataUp {
	m := make([]byte, 0)
	m = append(m, mhdr.New(mtype, major).ByteArray()...)
	m = append(m, fhdr.NewUp(devaddr, fcnt, fopts, adr,
		ack, adrackreq, foptslen).ByteArray()...)
	if fport != nil {
		m = append(m, *fport)
		m = append(m, frmpayload...)
	}
	m = append(m, mic...)
	return dataUp(m)
}

func ParseUp(p, nwkskey, appskey []byte) (DataUp, error) {
	temp := make([]byte, len(p))
	copy(temp, p)
	if len(p) < 12 {
		return nil, fmt.Errorf("payload should have length at least 12: %d", len(p))
	}
	fheader := fhdr.ParseUp(p)

	if len(p) > 1+7+int(fheader.FOptsLen())+4 { //have Fport
		fport := uint8(p[1+7+fheader.FOptsLen()])
		start := int(1 + 7 + fheader.FOptsLen() + 1)
		finish := len(p) - 4

		if fport == 0 {
			if nwkskey != nil {
				payload := encrypt(p[start:finish], true, fheader.DevAddr(), uint32(fheader.FCnt()), nwkskey)
				copy(temp[start:finish], payload)
			} else {
				return nil, fmt.Errorf("nwkskey must have a valid key")
			}
		} else {
			if appskey != nil {
				payload := encrypt(p[start:finish], true, fheader.DevAddr(), uint32(fheader.FCnt()), appskey)
				copy(temp[start:finish], payload)
			} else {
				return nil, fmt.Errorf("appskey must have a valid key")
			}
		}
	}
	return dataUp(temp), nil
}
