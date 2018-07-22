package macpayload

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/tkiraly/lorawan/commands"
	"github.com/tkiraly/lorawan/fhdr"
	"github.com/tkiraly/lorawan/mhdr"
)

type DataDown interface {
	MType() mhdr.MType
	Major() mhdr.MajorVersion
	fhdr.FHDRDown
	FPort() *uint8
	FRMPayload() []byte
	MIC() []byte
}

type dataDown []byte

func (m dataDown) ByteArray() []byte {
	return m
}

func (m dataDown) MType() mhdr.MType {
	v := mhdr.Parse(m[0])
	return v.MType()
}

func (m dataDown) Major() mhdr.MajorVersion {
	v := mhdr.Parse(m[0])
	return v.Major()
}

func (d dataDown) FPort() *uint8 {
	f := fhdr.ParseDown(d[1:])
	i := 1 + 7 + int(f.FOptsLen()) + 4
	if len(d) > i {
		return &d[1+7+f.FOptsLen()]
	}
	return nil
}

func (d dataDown) FRMPayload() []byte {
	f := fhdr.ParseDown(d[1:])
	i := 1 + 7 + int(f.FOptsLen()) + 4
	if len(d) > i {
		return d[1+7+f.FOptsLen()+1 : len(d)-4]
	}
	return nil
}

func (d dataDown) MIC() []byte {
	return d[len(d)-4:]
}

func (du dataDown) String() string {
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
func (du dataDown) DevAddr() []byte {
	d := fhdr.ParseDown(du[1:])
	return d.DevAddr()
}
func (du dataDown) FCnt() uint16 {
	d := fhdr.ParseDown(du[1:])
	return d.FCnt()
}
func (du dataDown) FOpts() []commands.Fopter {
	d := fhdr.ParseDown(du[1:])
	return d.FOpts()
}
func (du dataDown) FPending() bool {
	d := fhdr.ParseDown(du[1:])
	return d.FPending()
}
func (du dataDown) ADR() bool {
	d := fhdr.ParseDown(du[1:])
	return d.ADR()
}
func (du dataDown) ACK() bool {
	d := fhdr.ParseDown(du[1:])
	return d.ACK()
}
func (du dataDown) FOptsLen() uint8 {
	d := fhdr.ParseDown(du[1:])
	return d.FOptsLen()
}

func NewDown(mtype mhdr.MType, major mhdr.MajorVersion, devaddr []byte,
	fcnt uint16, fopts []commands.Fopter, adr,
	ack, fpending bool, foptslen uint8, fport *uint8,
	frmpayload, mic []byte) DataDown {
	m := make([]byte, 0)
	m = append(m, mhdr.New(mtype, major).ByteArray()...)
	m = append(m, fhdr.NewDown(devaddr, fcnt, fopts, adr,
		ack, fpending, foptslen).ByteArray()...)
	if fport != nil {
		m = append(m, *fport)
		m = append(m, frmpayload...)
	}
	m = append(m, mic...)
	return dataDown(m)
}

func ParseDown(p, nwkskey, appskey []byte) (DataDown, error) {
	temp := make([]byte, len(p))
	copy(temp, p)
	if len(p) < 12 {
		return nil, fmt.Errorf("payload should have length at least 12: %d", len(p))
	}
	fheader := fhdr.ParseDown(p)

	if len(p) > 1+7+int(fheader.FOptsLen())+4 { //have Fport
		fport := uint8(p[1+7+fheader.FOptsLen()])
		start := int(1 + 7 + fheader.FOptsLen() + 1)
		finish := len(p) - 4

		if fport == 0 {
			if nwkskey != nil {
				payload := encrypt(p[start:finish], false, fheader.DevAddr(), uint32(fheader.FCnt()), nwkskey)
				copy(temp[start:finish], payload)
			} else {
				return nil, fmt.Errorf("nwkskey must have a valid key")
			}
		} else {
			if appskey != nil {
				payload := encrypt(p[start:finish], false, fheader.DevAddr(), uint32(fheader.FCnt()), appskey)
				copy(temp[start:finish], payload)
			} else {
				return nil, fmt.Errorf("appskey must have a valid key")
			}
		}
	}
	return dataDown(temp), nil
}
