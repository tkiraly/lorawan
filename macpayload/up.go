package macpayload

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/tkiraly/lorawan/commands"
	"github.com/tkiraly/lorawan/fhdr"
	"github.com/tkiraly/lorawan/mhdr"
	"github.com/tkiraly/lorawan/mic"
)

type DataUp interface {
	MType() mhdr.MType
	Major() mhdr.MajorVersion
	fhdr.FHDRUp
	FPort() *uint8
	FRMPayload() []byte
	MIC() []byte
}
type dataUp struct {
	bytes, nwkskey, appskey []byte
}

func (du dataUp) ByteArray() []byte {
	temp := make([]byte, len(du.bytes))
	copy(temp, du.bytes)
	if du.FPort() != nil {
		key := du.appskey
		if *du.FPort() == 0 {
			key = du.nwkskey
		}
		fpay := endataMessage(du.FRMPayload(), true, du.DevAddr(), uint32(du.FCnt()), key)
		start := int(1 + 7 + du.FOptsLen() + 1)
		finish := len(du.bytes) - 4
		copy(temp[start:finish], fpay)
	}
	return temp
}

func (m dataUp) MType() mhdr.MType {
	v := mhdr.Parse(m.bytes[0])
	return v.MType()
}

func (m dataUp) Major() mhdr.MajorVersion {
	v := mhdr.Parse(m.bytes[0])
	return v.Major()
}

func (d dataUp) FPort() *uint8 {
	f := fhdr.ParseDown(d.bytes[1:])
	i := 1 + 7 + int(f.FOptsLen()) + 4
	if len(d.bytes) > i {
		return &d.bytes[1+7+f.FOptsLen()]
	}
	return nil
}

func (d dataUp) FRMPayload() []byte {
	f := fhdr.ParseDown(d.bytes[1:])
	i := 1 + 7 + int(f.FOptsLen()) + 4
	if len(d.bytes) > i {
		return d.bytes[1+7+f.FOptsLen()+1 : len(d.bytes)-4]
	}
	return nil
}

func (d dataUp) MIC() []byte {
	return d.bytes[len(d.bytes)-4:]
}

func (du dataUp) String() string {
	fport := "none"
	if du.FPort() != nil {
		fport = strconv.Itoa(int(*du.FPort()))
	}
	return fmt.Sprintf("DataUp! %s; FHDRUp: %s; FPort: %s; FRMPayload: %s; MIC: %s",
		mhdr.Parse(du.bytes[0]),
		fhdr.ParseUp(du.bytes[1:]),
		fport,
		strings.ToUpper(hex.EncodeToString(du.FRMPayload())),
		strings.ToUpper(hex.EncodeToString(du.MIC())))
}
func (du dataUp) DevAddr() []byte {
	d := fhdr.ParseUp(du.bytes[1:])
	return d.DevAddr()
}
func (du dataUp) FCnt() uint16 {
	d := fhdr.ParseUp(du.bytes[1:])
	return d.FCnt()
}
func (du dataUp) FOpts() []commands.Fopter {
	d := fhdr.ParseUp(du.bytes[1:])
	return d.FOpts()
}
func (du dataUp) ADRACKReq() bool {
	d := fhdr.ParseUp(du.bytes[1:])
	return d.ADRACKReq()
}
func (du dataUp) ADR() bool {
	d := fhdr.ParseUp(du.bytes[1:])
	return d.ADR()
}
func (du dataUp) ACK() bool {
	d := fhdr.ParseUp(du.bytes[1:])
	return d.ACK()
}
func (du dataUp) FOptsLen() uint8 {
	d := fhdr.ParseUp(du.bytes[1:])
	return d.FOptsLen()
}

func NewUp(mtype mhdr.MType, major mhdr.MajorVersion, devaddr []byte,
	fcnt uint16, fopts []commands.Fopter, adr,
	ack, adrackreq bool, foptslen uint8, fport *uint8,
	frmpayload, nwkskey, appskey []byte) DataUp {
	m := make([]byte, 0)
	m = append(m, mhdr.New(mtype, major).ByteArray()...)
	m = append(m, fhdr.NewUp(devaddr, fcnt, fopts, adr,
		ack, adrackreq, foptslen).ByteArray()...)
	fstart := 0
	ffinish := 0
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
		fpay := endataMessage(frmpayload, true, devaddr, uint32(fcnt), key)
		m = append(m, fpay...)
	}
	m = append(m, make([]byte, 4)...)
	mi, err := mic.Calculate(m, nwkskey)

	if err != nil {
		panic(err)
	}
	copy(m[len(m)-4:], mi)

	if fport != nil {
		key := appskey
		if *fport == 0 {
			key = nwkskey
		}
		fpay := endataMessage(m[fstart:ffinish], true, devaddr, uint32(fcnt), key)
		copy(m[fstart:ffinish], fpay)
	}

	return dataUp{bytes: m, nwkskey: nwkskey, appskey: appskey}
}

func ParseUp(p, nwkskey, appskey []byte) (DataUp, error) {
	temp := make([]byte, len(p))
	copy(temp, p)
	if len(p) < 12 {
		return nil, fmt.Errorf("payload should have length at least 12: %d", len(p))
	}
	du := dataUp{bytes: temp, nwkskey: nwkskey, appskey: appskey}
	if du.FPort() != nil {
		key := appskey
		if *du.FPort() == 0 {
			key = nwkskey
		}
		fpay := endataMessage(du.FRMPayload(), true, du.DevAddr(), uint32(du.FCnt()), key)
		start := int(1 + 7 + du.FOptsLen() + 1)
		finish := len(du.bytes) - 4
		copy(du.bytes[start:finish], fpay)
	}
	return du, nil
}

func endataMessage(payload []byte, updir bool, devaddr []byte, cnt uint32, key []byte) []byte {
	var dir byte = 0x01
	if updir {
		dir = 0x00
	}

	/*
	   this is an example how Block a is formed according to Lorawan specs
	   Block_A[0]  = 0x01;
	   Block_A[1]  = 0x00;
	   Block_A[2]  = 0x00;
	   Block_A[3]  = 0x00;
	   Block_A[4]  = 0x00;
	   Block_A[5]  = 0;        // 0 for uplink frames 1 for downlink frames;
	   Block_A[6]  = dev_addr[3]; // LSB devAddr 4 bytes
	   Block_A[7]  = dev_addr[2];  // ..
	   Block_A[8]  = dev_addr[1];  // ..
	   Block_A[9]  = dev_addr[0];  // MSB
	   Block_A[10] = sequence_counter & 0xff;  // LSB framecounter
	   Block_A[11] = (sequence_counter >> 8) & 0xff;  // MSB framecounter
	   Block_A[12] = (sequence_counter >> 16) & 0xff;     // Frame counter upper Bytes
	   Block_A[13] = (sequence_counter >> 24) & 0xff;
	   Block_A[14] = 0x00;
	*/

	block := []byte{0x01, 0x00, 0x00, 0x00,
		0x00, dir, devaddr[3], devaddr[2],
		devaddr[1], devaddr[0], byte(cnt & 0xff), byte((cnt >> 8) & 0xff),
		byte((cnt >> 16) & 0xff), byte((cnt >> 24) & 0xff), 0x00, 0x00}

	cipher, _ := aes.NewCipher(key)
	r := make([]byte, len(payload))
	bufferIndex := 0
	size := len(payload)
	ctr := 1
	S := make([]byte, 16)

	for size >= 16 {
		block[15] = byte(ctr & 0xFF)
		ctr += 1
		cipher.Encrypt(S, block)

		for i := 0; i < 16; i++ {
			r[bufferIndex+i] = payload[bufferIndex+i] ^ S[i]
		}
		size -= 16
		bufferIndex += 16
	}
	if size > 0 {
		block[15] = byte(ctr & 0xFF)
		ctr += 1

		cipher.Encrypt(S, block)
		for i := 0; i < size; i++ {
			r[bufferIndex+i] = payload[bufferIndex+i] ^ S[i]
		}
	}
	return r
}

func dedataMessage(payload []byte, updir bool, devaddr []byte, cnt uint32, key []byte) []byte {
	var dir byte = 0x01
	if updir {
		dir = 0x00
	}

	/*
	   this is an example how Block a is formed according to Lorawan specs
	   Block_A[0]  = 0x01;
	   Block_A[1]  = 0x00;
	   Block_A[2]  = 0x00;
	   Block_A[3]  = 0x00;
	   Block_A[4]  = 0x00;
	   Block_A[5]  = 0;        // 0 for uplink frames 1 for downlink frames;
	   Block_A[6]  = dev_addr[3]; // LSB devAddr 4 bytes
	   Block_A[7]  = dev_addr[2];  // ..
	   Block_A[8]  = dev_addr[1];  // ..
	   Block_A[9]  = dev_addr[0];  // MSB
	   Block_A[10] = sequence_counter & 0xff;  // LSB framecounter
	   Block_A[11] = (sequence_counter >> 8) & 0xff;  // MSB framecounter
	   Block_A[12] = (sequence_counter >> 16) & 0xff;     // Frame counter upper Bytes
	   Block_A[13] = (sequence_counter >> 24) & 0xff;
	   Block_A[14] = 0x00;
	*/

	block := []byte{0x01, 0x00, 0x00, 0x00,
		0x00, dir, devaddr[3], devaddr[2],
		devaddr[1], devaddr[0], byte(cnt & 0xff), byte((cnt >> 8) & 0xff),
		byte((cnt >> 16) & 0xff), byte((cnt >> 24) & 0xff), 0x00, 0x00}

	cipher, _ := aes.NewCipher(key)
	r := make([]byte, len(payload))
	bufferIndex := 0
	size := len(payload)
	ctr := 1
	S := make([]byte, 16)

	for size >= 16 {
		block[15] = byte(ctr & 0xFF)
		ctr += 1
		cipher.Decrypt(S, block)

		for i := 0; i < 16; i++ {
			r[bufferIndex+i] = payload[bufferIndex+i] ^ S[i]
		}
		size -= 16
		bufferIndex += 16
	}
	if size > 0 {
		block[15] = byte(ctr & 0xFF)
		ctr += 1

		cipher.Decrypt(S, block)
		for i := 0; i < size; i++ {
			r[bufferIndex+i] = payload[bufferIndex+i] ^ S[i]
		}
	}
	return r
}
