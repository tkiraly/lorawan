package joinaccept

import (
	"crypto/aes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/tkiraly/lorawan/mhdr"
	"github.com/tkiraly/lorawan/mic"
	"github.com/tkiraly/lorawan/util"
)

//JoinAccept serve the methods which can be used on a Join Accept message
type JoinAccept interface {
	mhdr.MHDR
	AppNonce() []byte
	NetID() []byte
	DevAddr() []byte
	RxDelay() byte
	CFlist() []uint32
	DlSettingsRX1DRoffset() byte
	DlSettingsRX2Datarate() byte
	MIC() []byte
}

type joinAccept struct {
	bytes, key []byte
}

//New create a new Join Accept
func New(Major mhdr.MajorVersion, appnonce, netid, devaddr []byte, rx1droffset, rx2datarate, rxdelay byte,
	cflist []uint32, key []byte) JoinAccept {
	ja := make([]byte, 17)
	ja[0] = mhdr.New(mhdr.JoinAcceptMessageType, mhdr.LoRaWANR1MajorVersion).ByteArray()[0]
	copy(ja[1:4], util.Bytereverse(appnonce))
	copy(ja[4:7], util.Bytereverse(netid))
	copy(ja[7:11], util.Bytereverse(devaddr))
	ja[11] = ((rx1droffset << 4) & rx2datarate) & 0x7F
	ja[12] = rxdelay
	for i := 0; i < len(cflist); i++ {
		ch := make([]byte, 4)
		binary.BigEndian.PutUint32(ch, cflist[i])
		ja = append(ja, util.Bytereverse(ch[1:])...)
	}
	if key != nil {
		m, err := mic.Calculate(ja, key)
		if err != nil {
			panic(err)
		}
		copy(ja[len(ja)-4:], m)
	}
	return joinAccept{bytes: ja, key: key}
}

//Parse converts a byte array to a Join Accept
func Parse(bb, key []byte) (JoinAccept, error) {
	if key == nil {
		if len(bb) != 13 && len(bb) != 29 {
			return nil, fmt.Errorf("payload should have length 13 or 29 but have: %d", len(bb))
		}
		return joinAccept{bytes: bb, key: key}, nil
	}
	if len(bb) != 17 && len(bb) != 33 {
		return nil, fmt.Errorf("payload should have length 17 or 33 but have: %d", len(bb))
	}
	b := make([]byte, len(bb))
	copy(b, bb)
	acc := joinAccept{bytes: b, key: key}
	acc.encrypt()
	return acc, nil
}

func (ja *joinAccept) encrypt() {
	if ja.key != nil {
		cipher, _ := aes.NewCipher(ja.key)
		i := 1
		if len(ja.bytes) > 17 {
			i++
		}
		r := make([]byte, 16*i)
		for index := 0; index < i; index++ {
			cipher.Encrypt(r[16*index:16*index+16],
				ja.bytes[16*index+1:16*index+17])
		}
		ja.bytes = []byte{ja.bytes[0]}
		ja.bytes = append(ja.bytes, r...)
	}
}

func (ja joinAccept) decrypt() []byte {
	if ja.key != nil {
		cipher, _ := aes.NewCipher(ja.key)
		i := 1
		if len(ja.bytes) > 17 {
			i++
		}
		r := make([]byte, 16*i)
		for index := 0; index < i; index++ {
			cipher.Decrypt(r[16*index:16*index+16],
				ja.bytes[16*index+1:16*index+17])
		}
		a := []byte{ja.bytes[0]}
		a = append(a, r...)
		return a
	}
	return nil
}

func (ja joinAccept) AppNonce() []byte {
	return util.Bytereverse(ja.bytes[1:4])
}
func (ja joinAccept) NetID() []byte {
	return util.Bytereverse(ja.bytes[4:7])
}
func (ja joinAccept) DevAddr() []byte {
	return util.Bytereverse(ja.bytes[7:11])
}
func (ja joinAccept) DlSettingsRX1DRoffset() byte {
	return (ja.bytes[11] >> 4) & 0x07
}
func (ja joinAccept) DlSettingsRX2Datarate() byte {
	return ja.bytes[11] & 0x0F
}
func (ja joinAccept) RxDelay() byte {
	return ja.bytes[12]
}
func (ja joinAccept) CFlist() []uint32 {
	if len(ja.bytes) > 17 {
		chs := make([]uint32, 5)
		temp := make([]byte, 4)
		for i := 0; i < 5; i++ {
			copy(temp, ja.bytes[13+(3*i):16+(3*i)])
			chs[i] = binary.LittleEndian.Uint32(temp)
		}
		return chs
	}
	return nil
}
func (ja joinAccept) MIC() []byte {
	if len(ja.bytes) > 17 {
		return ja.bytes[29:]
	}
	return ja.bytes[13:]
}

func (ja joinAccept) MType() mhdr.MType {
	m := mhdr.Parse(ja.bytes[0])
	return m.MType()
}

func (ja joinAccept) Major() mhdr.MajorVersion {
	m := mhdr.Parse(ja.bytes[0])
	return m.Major()
}

func (ja joinAccept) String() string {
	return fmt.Sprintf("%s: AppNonce: %s; NetID: %s; DevAddr: %s; RX1DRoffset: %s; RX2Datarate: %s; RxDelay: %s; MIC: %s; CFList: %v\n",
		ja.MType(),
		strings.ToUpper(hex.EncodeToString(ja.AppNonce())),
		strings.ToUpper(hex.EncodeToString(ja.NetID())),
		strings.ToUpper(hex.EncodeToString(ja.DevAddr())),
		strings.ToUpper(hex.EncodeToString([]byte{ja.DlSettingsRX1DRoffset()})),
		strings.ToUpper(hex.EncodeToString([]byte{ja.DlSettingsRX2Datarate()})),
		strings.ToUpper(hex.EncodeToString([]byte{ja.RxDelay()})),
		strings.ToUpper(hex.EncodeToString(ja.MIC())),
		ja.CFlist())
}

func (ja joinAccept) ByteArray() []byte {

	return ja.decrypt()
}
