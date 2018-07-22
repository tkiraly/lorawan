package joinaccept

import (
	"crypto/aes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/tkiraly/lorawan/mhdr"
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

type joinAccept []byte

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
	ja = encrypt(ja, key)
	return joinAccept(ja)
}

//Parse converts a byte array to a Join Accept
func Parse(bb []byte) (JoinAccept, error) {
	b := make([]byte, len(bb))
	copy(b, bb)
	if len(b) != 17 && len(b) != 33 {
		return nil, fmt.Errorf("payload should have length 17 or 33 but have: %d", len(b))
	}
	return joinAccept(b), nil
}

func ParseEncrypted(bb, key []byte) (JoinAccept, error) {
	p := encrypt(bb, key)
	return Parse(p)
}

func encrypt(payload, key []byte) []byte {
	if len(payload) > 17 {
		cipher, _ := aes.NewCipher(key)
		r := make([]byte, 32)
		cipher.Encrypt(r[0:16], payload[1:17])
		cipher.Encrypt(r[16:32], payload[17:33])
		return append([]byte{payload[0]}, r...)
	}
	cipher, _ := aes.NewCipher(key)
	r := make([]byte, 16)
	cipher.Encrypt(r[0:16], payload[1:17])
	return append([]byte{payload[0]}, r...)
}

func (ja joinAccept) AppNonce() []byte {
	return util.Bytereverse(ja[1:4])
}
func (ja joinAccept) NetID() []byte {
	return util.Bytereverse(ja[4:7])
}
func (ja joinAccept) DevAddr() []byte {
	return util.Bytereverse(ja[7:11])
}
func (ja joinAccept) DlSettingsRX1DRoffset() byte {
	return (ja[11] >> 4) & 0x07
}
func (ja joinAccept) DlSettingsRX2Datarate() byte {
	return ja[11] & 0x0F
}
func (ja joinAccept) RxDelay() byte {
	return ja[12]
}
func (ja joinAccept) CFlist() []uint32 {
	if len(ja) > 17 {
		chs := make([]uint32, 5)
		temp := make([]byte, 4)
		for i := 0; i < 5; i++ {
			copy(temp, ja[13+(3*i):16+(3*i)])
			chs[i] = binary.LittleEndian.Uint32(temp)
		}
		return chs
	}
	return nil
}
func (ja joinAccept) MIC() []byte {
	if len(ja) > 17 {
		return ja[29:]
	}
	return ja[13:]
}

func (ja joinAccept) MType() mhdr.MType {
	m := mhdr.Parse(ja[0])
	return m.MType()
}

func (ja joinAccept) Major() mhdr.MajorVersion {
	m := mhdr.Parse(ja[0])
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
	return ja
}
