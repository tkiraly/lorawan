package mhdr

import (
	"reflect"
	"testing"
)

func Test_mHDR_MType(t *testing.T) {
	tests := []struct {
		name string
		m    MHDR
		want MType
	}{
		{"basic", New(ConfirmedDataDownMessageType, LoRaWANR1MajorVersion), ConfirmedDataDownMessageType},
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
		m    MHDR
		want MajorVersion
	}{
		{"basic", New(ConfirmedDataDownMessageType, LoRaWANR1MajorVersion), LoRaWANR1MajorVersion},
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
		m    MHDR
		want []byte
	}{
		{"basic", New(ConfirmedDataDownMessageType, LoRaWANR1MajorVersion), []byte{0xA0}},
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
		want MHDR
	}{
		{"basic", args{b: 0xA0}, New(ConfirmedDataDownMessageType, LoRaWANR1MajorVersion)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		mtype MType
		major MajorVersion
	}
	tests := []struct {
		name string
		args args
		want MHDR
	}{
		{"basic", args{mtype: ConfirmedDataDownMessageType, major: LoRaWANR1MajorVersion},
			New(ConfirmedDataDownMessageType, LoRaWANR1MajorVersion)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.mtype, tt.args.major); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mHDR_String(t *testing.T) {
	tests := []struct {
		name string
		m    mHDR
		want string
	}{
		{"basic",
			mHDR(0x40),
			"MType: UnconfirmedDataUpMessageType, Major: LoRaWANR1MajorVersion;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("mHDR.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
