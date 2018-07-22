package mhdr_test

import (
	"fmt"

	"github.com/tkiraly/lorawan/mhdr"
)

func ExampleMType() {
	m := mhdr.New(mhdr.ConfirmedDataDownMessageType, mhdr.LoRaWANR1MajorVersion)

	fmt.Printf("%s", m.MType())
	// Output: ConfirmedDataDownMessageType
}
