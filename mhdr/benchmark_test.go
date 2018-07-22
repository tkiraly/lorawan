package mhdr_test

import (
	"testing"

	"github.com/tkiraly/lorawan/mhdr"
)

func Benchmark_mHDR_MType(b *testing.B) {
	m := mhdr.New(mhdr.ConfirmedDataDownMessageType, mhdr.LoRaWANR1MajorVersion)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.MType()
	}
}

func Benchmark_mHDR_Major(b *testing.B) {
	m := mhdr.New(mhdr.ConfirmedDataDownMessageType, mhdr.LoRaWANR1MajorVersion)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Major()
	}
}
