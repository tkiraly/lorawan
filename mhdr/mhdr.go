/*
Package mhdr contains LoRaWAN MAC Header related stuff.

The MAC header specifies the message type (MType) and according to which major version
(Major) of the frame format of the LoRaWAN layer specification the frame has been
encoded.
*/
package mhdr

import (
	"fmt"
)

//go:generate stringer -type=MType
//go:generate stringer -type=MajorVersion

//MType represents the message type
type MType byte

//The LoRaWAN distinguishes between six different MAC message types: join request, join
//accept, unconfirmed data up/down, and confirmed data up/down.
const (
	JoinRequestMessageType MType = iota
	JoinAcceptMessageType
	UnconfirmedDataUpMessageType
	UnconfirmedDataDownMessageType
	ConfirmedDataUpMessageType
	ConfirmedDataDownMessageType
	RFUMessageType
	ProprietaryMessageType
)

// MajorVersion represents the major version
type MajorVersion byte

// represents the Major version
const (
	LoRaWANR1MajorVersion MajorVersion = iota
)

// Parse a byte into a MAC header.
func Parse(bb byte) MHDR {
	b := bb
	return mHDR(b)
}

// New returns a new MAC header.
func New(mtype MType, major MajorVersion) MHDR {
	return mHDR((byte(mtype) << 5) | (byte(major) << 0))
}

// MHDR is accessed returns an os.FileInfo for the FileHeader.
type MHDR interface {
	//String returns a MAC header into a string
	String() string
	//ByteArray returns a MAC header into a byte array
	ByteArray() []byte
	//MType returns message type
	MType() MType
	//Major returns major version
	Major() MajorVersion
}

type mHDR byte

func (m mHDR) MType() MType {
	return (MType)((m >> 5) & 0x07)
}

func (m mHDR) Major() MajorVersion {
	return (MajorVersion)(m & 0x03)
}

func (m mHDR) String() string {
	return fmt.Sprintf("MType: %s; Major: %s;", m.MType(), m.Major())
}

func (m mHDR) ByteArray() []byte {
	return []byte{byte(m)}
}
