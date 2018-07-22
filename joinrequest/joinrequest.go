package joinrequest

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/jacobsa/crypto/cmac"
	"github.com/tkiraly/lorawan/mhdr"
	"github.com/tkiraly/lorawan/util"
)

//JoinRequest serve the methods which can be used on a Join Request message
type JoinRequest interface {
	mhdr.MHDR
	AppEUI() []byte
	DevEUI() []byte
	DevNonce() []byte
	MIC() []byte
}

//New create a new Join Request
func New(Major mhdr.MajorVersion, appeui, deveui, appkey, devnonce []byte) (JoinRequest, error) {
	jr := make([]byte, 23)
	jr[0] = mhdr.New(mhdr.JoinRequestMessageType, Major).ByteArray()[0]
	copy(jr[1:9], util.Bytereverse(appeui))
	copy(jr[9:17], util.Bytereverse(deveui))

	if devnonce == nil {
		rand.Seed(time.Now().UTC().Unix())
		devnonce = []byte{
			byte(rand.Uint32()),
			byte(rand.Uint32()),
		}
	}
	copy(jr[17:19], util.Bytereverse(devnonce))
	Jr := joinRequest(jr)
	m, err := calculatemic(Jr.ByteArray(), appkey)
	if err != nil {
		return nil, err
	}
	copy(jr[19:23], m)
	return Jr, nil
}

//Parse converts a byte array to a Join Request
func Parse(bb []byte) (JoinRequest, error) {
	b := make([]byte, len(bb))
	copy(b, bb)
	if len(b) != 23 {
		return nil, fmt.Errorf("payload should have length 23 but have: %d", len(b))
	}

	return joinRequest(b), nil
}

func calculatemic(payload, key []byte) ([]byte, error) {
	mhd := mhdr.Parse(payload[0])
	B := []byte{0x49, 0x00, 0x00, 0x00, 0x00}
	switch mhd.MType() {
	case mhdr.JoinRequestMessageType, mhdr.JoinAcceptMessageType:
		B = payload[:len(payload)-4]
	}
	result, err := cmac.New(key)
	if err != nil {
		return nil, err
	}
	result.Write(B)
	rr := result.Sum([]byte{})
	return rr[:4], nil
}

type joinRequest []byte

func (jr joinRequest) AppEUI() []byte {
	return util.Bytereverse(jr[1:9])
}

func (jr joinRequest) DevEUI() []byte {
	return util.Bytereverse(jr[9:17])
}

func (jr joinRequest) DevNonce() []byte {
	return util.Bytereverse(jr[17:19])
}

func (jr joinRequest) MIC() []byte {
	return jr[19:23]
}

func (jr joinRequest) MType() mhdr.MType {
	m := mhdr.Parse(jr[0])
	return m.MType()
}

func (jr joinRequest) Major() mhdr.MajorVersion {
	m := mhdr.Parse(jr[0])
	return m.Major()
}

func (jr joinRequest) String() string {
	return fmt.Sprintf("%s: AppEUI: %s; DevEUI: %s; DevNonce: %s; MIC: %s\n",
		jr.MType(),
		strings.ToUpper(hex.EncodeToString(jr.AppEUI())),
		strings.ToUpper(hex.EncodeToString(jr.DevEUI())),
		strings.ToUpper(hex.EncodeToString(jr.DevNonce())),
		strings.ToUpper(hex.EncodeToString(jr.MIC())))
}

func (jr joinRequest) ByteArray() []byte {
	return jr
}
