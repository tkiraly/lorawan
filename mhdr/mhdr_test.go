package mhdr_test

import (
	"reflect"
	"testing"

	"github.com/tkiraly/lorawan/mhdr"
)

func Test_mHDR_MType(t *testing.T) {
	tests := []struct {
		name string
		m    mhdr.MHDR
		want mhdr.MType
	}{
		{"basic", mhdr.New(mhdr.ConfirmedDataDownMessageType, mhdr.LoRaWANR1MajorVersion), mhdr.ConfirmedDataDownMessageType},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.MType(); got != tt.want {
				t.Errorf("mHDR.MType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mHDR_Major(t *testing.T) {
	tests := []struct {
		name string
		m    mhdr.MHDR
		want mhdr.MajorVersion
	}{
		{"basic", mhdr.New(mhdr.ConfirmedDataDownMessageType, mhdr.LoRaWANR1MajorVersion), mhdr.LoRaWANR1MajorVersion},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Major(); got != tt.want {
				t.Errorf("mHDR.Major() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mHDR_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		m    mhdr.MHDR
		want []byte
	}{
		{"basic", mhdr.New(mhdr.ConfirmedDataDownMessageType, mhdr.LoRaWANR1MajorVersion), []byte{0xA0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mHDR.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want mhdr.MHDR
	}{
		{"basic", args{b: 0xA0}, mhdr.New(mhdr.ConfirmedDataDownMessageType, mhdr.LoRaWANR1MajorVersion)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mhdr.Parse(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		mtype mhdr.MType
		major mhdr.MajorVersion
	}
	tests := []struct {
		name string
		args args
		want mhdr.MHDR
	}{
		{"basic", args{mtype: mhdr.ConfirmedDataDownMessageType, major: mhdr.LoRaWANR1MajorVersion},
			mhdr.New(mhdr.ConfirmedDataDownMessageType, mhdr.LoRaWANR1MajorVersion)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mhdr.New(tt.args.mtype, tt.args.major); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
