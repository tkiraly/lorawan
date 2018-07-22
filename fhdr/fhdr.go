package fhdr

import (
	"github.com/tkiraly/lorawan/commands"
)

type FHDR interface {
	//String returns a MAC header into a string
	String() string
	//ByteArray returns a MAC header into a byte array
	ByteArray() []byte
	DevAddr() []byte
	FCnt() uint16
	FOpts() []commands.Fopter
}

type FCtrl interface {
	ADR() bool
	ACK() bool
	FOptsLen() uint8
}
