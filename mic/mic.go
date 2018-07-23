package mic

import (
	"encoding/binary"

	"github.com/jacobsa/crypto/cmac"
	"github.com/tkiraly/lorawan/fhdr"
	"github.com/tkiraly/lorawan/mhdr"
	"github.com/tkiraly/lorawan/util"
)

func Calculate(payload, key []byte) ([]byte, error) {
	mhd := mhdr.Parse(payload[0])
	B := []byte{0x49, 0x00, 0x00, 0x00, 0x00}
	switch mhd.MType() {
	case mhdr.UnconfirmedDataDownMessageType, mhdr.ConfirmedDataDownMessageType:
		fheader := fhdr.ParseDown(payload[1:])
		B = append(B, 1)
		B = append(B, util.Bytereverse(fheader.DevAddr())...)
		fcnt := make([]byte, 4)
		binary.LittleEndian.PutUint16(fcnt, fheader.FCnt())
		B = append(B, fcnt...)
		B = append(B, 0)
		B = append(B, byte(len(payload)-4))
		B = append(B, payload[:len(payload)-4]...)

	case mhdr.UnconfirmedDataUpMessageType, mhdr.ConfirmedDataUpMessageType:
		fheader := fhdr.ParseUp(payload[1:])
		B = append(B, 0)
		B = append(B, util.Bytereverse(fheader.DevAddr())...)
		fcnt := make([]byte, 4)
		binary.LittleEndian.PutUint16(fcnt, fheader.FCnt())
		B = append(B, fcnt...)
		B = append(B, 0)
		B = append(B, byte(len(payload)-4))
		B = append(B, payload[:len(payload)-4]...)
	case mhdr.JoinRequestMessageType, mhdr.JoinAcceptMessageType:
		B = payload[:len(payload)-4]
	}
	result, err := cmac.New(key)
	if err != nil {
		return nil, err
	}
	_, err = result.Write(B)
	if err != nil {
		panic(err)
	}
	rr := result.Sum([]byte{})
	return rr[:4], nil
}
