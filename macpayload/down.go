package macpayload

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/tkiraly/lorawan/commands"
	"github.com/tkiraly/lorawan/fhdr"
	"github.com/tkiraly/lorawan/mhdr"
	"github.com/tkiraly/lorawan/mic"
)

type DataDown interface {
	MType() mhdr.MType
	Major() mhdr.MajorVersion
	fhdr.FHDRDown
	FPort() *uint8
	FRMPayload() []byte
	MIC() []byte
}

type dataDown struct {
	bytes, nwkskey, appskey []byte
}

func (m dataDown) ByteArray() []byte {
	temp := make([]byte, len(m.bytes))
	copy(temp, m.bytes)
	if m.FPort() != nil {
		key := m.appskey
		if *m.FPort() == 0 {
			key = m.nwkskey
		}
		fpay := endataMessage(m.FRMPayload(), false, m.DevAddr(), uint32(m.FCnt()), key)
		start := int(1 + 7 + m.FOptsLen() + 1)
		finish := len(m.bytes) - 4
		copy(temp[start:finish], fpay)
	}
	return temp
}

func (m dataDown) MType() mhdr.MType {
	v := mhdr.Parse(m.bytes[0])
	return v.MType()
}

func (m dataDown) Major() mhdr.MajorVersion {
	v := mhdr.Parse(m.bytes[0])
	return v.Major()
}

func (d dataDown) FPort() *uint8 {
	f := fhdr.ParseDown(d.bytes[1:])
	i := 1 + 7 + int(f.FOptsLen()) + 4
	if len(d.bytes) > i {
		return &d.bytes[1+7+f.FOptsLen()]
	}
	return nil
}

func (d dataDown) FRMPayload() []byte {
	f := fhdr.ParseDown(d.bytes[1:])
	i := 1 + 7 + int(f.FOptsLen()) + 4
	if len(d.bytes) > i {
		return d.bytes[1+7+f.FOptsLen()+1 : len(d.bytes)-4]
	}
	return nil
}

func (d dataDown) MIC() []byte {
	return d.bytes[len(d.bytes)-4:]
}

func (du dataDown) String() string {
	fport := "none"
	if du.FPort() != nil {
		fport = strconv.Itoa(int(*du.FPort()))
	}
	return fmt.Sprintf("DataDown: MHDR: %s; FHDRUp: %s; FPort: %s; FRMPayload: %s; MIC: %s",
		mhdr.Parse(du.bytes[0]),
		fhdr.ParseDown(du.bytes[1:]),
		fport,
		strings.ToUpper(hex.EncodeToString(du.FRMPayload())),
		strings.ToUpper(hex.EncodeToString(du.MIC())))
}
func (du dataDown) DevAddr() []byte {
	d := fhdr.ParseDown(du.bytes[1:])
	return d.DevAddr()
}
func (du dataDown) FCnt() uint16 {
	d := fhdr.ParseDown(du.bytes[1:])
	return d.FCnt()
}
func (du dataDown) FOpts() []commands.Fopter {
	d := fhdr.ParseDown(du.bytes[1:])
	return d.FOpts()
}
func (du dataDown) FPending() bool {
	d := fhdr.ParseDown(du.bytes[1:])
	return d.FPending()
}
func (du dataDown) ADR() bool {
	d := fhdr.ParseDown(du.bytes[1:])
	return d.ADR()
}
func (du dataDown) ACK() bool {
	d := fhdr.ParseDown(du.bytes[1:])
	return d.ACK()
}
func (du dataDown) FOptsLen() uint8 {
	d := fhdr.ParseDown(du.bytes[1:])
	return d.FOptsLen()
}

func NewDown(mtype mhdr.MType, major mhdr.MajorVersion, devaddr []byte,
	fcnt uint16, fopts []commands.Fopter, adr,
	ack, fpending bool, foptslen uint8, fport *uint8,
	frmpayload, nwkskey, appskey []byte) DataDown {
	m := make([]byte, 0)
	m = append(m, mhdr.New(mtype, major).ByteArray()...)
	m = append(m, fhdr.NewDown(devaddr, fcnt, fopts, adr,
		ack, fpending, foptslen).ByteArray()...)
	fstart := 0
	ffinish := 0

	fmt.Println(hex.EncodeToString(m))
	if fport != nil {
		//WARNING: MIC is calculated with encrypted FRMPayload
		//but in the end FRMPayload must return decrypted value
		m = append(m, *fport)

		key := appskey
		if *fport == 0 {
			key = nwkskey
		}
		fstart = len(m)
		ffinish = fstart + len(frmpayload)
		fpay := endataMessage(frmpayload, false, devaddr, uint32(fcnt), key)
		fmt.Println(hex.EncodeToString(fpay))
		m = append(m, fpay...)
	}
	m = append(m, make([]byte, 4)...)
	mi, err := mic.Calculate(m, nwkskey)

	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(m))
	fmt.Println(hex.EncodeToString(mi))
	copy(m[len(m)-4:], mi)

	if fport != nil {
		key := appskey
		if *fport == 0 {
			key = nwkskey
		}
		fpay := endataMessage(m[fstart:ffinish], false, devaddr, uint32(fcnt), key)
		copy(m[fstart:ffinish], fpay)
	}

	return dataDown{bytes: m, nwkskey: nwkskey, appskey: appskey}

}

func ParseDown(p, nwkskey, appskey []byte) (DataDown, error) {
	temp := make([]byte, len(p))
	copy(temp, p)
	if len(p) < 12 {
		return nil, fmt.Errorf("payload should have length at least 12: %d", len(p))
	}
	du := dataDown{bytes: temp, nwkskey: nwkskey, appskey: appskey}
	if du.FPort() != nil {
		key := appskey
		if *du.FPort() == 0 {
			key = nwkskey
		}
		fpay := endataMessage(du.FRMPayload(), false, du.DevAddr(), uint32(du.FCnt()), key)
		start := int(1 + 7 + du.FOptsLen() + 1)
		finish := len(du.bytes) - 4
		copy(du.bytes[start:finish], fpay)
	}
	return du, nil
}
