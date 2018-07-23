package fhdr

import (
	"reflect"
	"testing"

	"github.com/tkiraly/lorawan/commands"
)

func Test_fHDRUp_DevAddr(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRUp
		want []byte
	}{
		{"basic",
			fHDRUp([]byte{0xbd, 0x1f, 0x52, 0x48, 0x80, 0x04, 0x00}),
			[]byte{0x48, 0x52, 0x1f, 0xbd},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.DevAddr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fHDRUp.DevAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRUp_FCnt(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRUp
		want uint16
	}{
		{"basic",
			fHDRUp([]byte{0xbd, 0x1f, 0x52, 0x48, 0x80, 0x04, 0x00}),
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.FCnt(); got != tt.want {
				t.Errorf("fHDRUp.FCnt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRUp_FOpts(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRUp
		want []commands.Fopter
	}{
		{"basic",
			fHDRUp([]byte{0xbd, 0x1f, 0x52, 0x48, 0x80, 0x04, 0x00}),
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.FOpts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fHDRUp.FOpts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRUp_ADRACKReq(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRUp
		want bool
	}{
		{"basic",
			fHDRUp([]byte{0xbd, 0x1f, 0x52, 0x48, 0x80, 0x04, 0x00}),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.ADRACKReq(); got != tt.want {
				t.Errorf("fHDRUp.ADRACKReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRUp_ADR(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRUp
		want bool
	}{
		{"basic",
			fHDRUp([]byte{0xbd, 0x1f, 0x52, 0x48, 0x80, 0x04, 0x00}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.ADR(); got != tt.want {
				t.Errorf("fHDRUp.ADR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRUp_ACK(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRUp
		want bool
	}{
		{"basic",
			fHDRUp([]byte{0xbd, 0x1f, 0x52, 0x48, 0x80, 0x04, 0x00}),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.ACK(); got != tt.want {
				t.Errorf("fHDRUp.ACK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRUp_FOptsLen(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRUp
		want uint8
	}{
		{"basic",
			fHDRUp([]byte{0xbd, 0x1f, 0x52, 0x48, 0x80, 0x04, 0x00}),
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.FOptsLen(); got != tt.want {
				t.Errorf("fHDRUp.FOptsLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUp(t *testing.T) {
	type args struct {
		devaddr   []byte
		fcnt      uint16
		fopts     []commands.Fopter
		adr       bool
		ack       bool
		adrackreq bool
		foptslen  uint8
	}
	tests := []struct {
		name string
		args args
		want FHDRUp
	}{
		{"basic",

			args{
				devaddr:   []byte{0x48, 0x52, 0x1f, 0xbd},
				fcnt:      4,
				fopts:     nil,
				adr:       true,
				ack:       false,
				adrackreq: false,
				foptslen:  0,
			},
			fHDRUp([]byte{0xbd, 0x1f, 0x52, 0x48, 0x80, 0x04, 0x00}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUp(tt.args.devaddr, tt.args.fcnt, tt.args.fopts, tt.args.adr, tt.args.ack, tt.args.adrackreq, tt.args.foptslen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRUp_String(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRUp
		want string
	}{
		{"basic",
			fHDRUp([]byte{0xbd, 0x1f, 0x52, 0x48, 0x80, 0x04, 0x00}),
			"DevAddr: 48521FBD; FCnt: 4; ADR: true; ADRACKReq: false; ACK: false; FOptsLen: 0, FOpts: [];",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.String(); got != tt.want {
				t.Errorf("fHDRUp.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseUp(t *testing.T) {
	type args struct {
		bb []byte
	}
	tests := []struct {
		name string
		args args
		want FHDRUp
	}{
		{"basic",
			args{bb: []byte{0xbd, 0x1f, 0x52, 0x48, 0x80, 0x04, 0x00}},
			fHDRUp([]byte{0xbd, 0x1f, 0x52, 0x48, 0x80, 0x04, 0x00}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseUp(tt.args.bb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRUp_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRUp
		want []byte
	}{
		{"basic",
			fHDRUp([]byte{0xbd, 0x1f, 0x52, 0x48, 0x80, 0x04, 0x00}),
			[]byte{0xbd, 0x1f, 0x52, 0x48, 0x80, 0x04, 0x00},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fHDRUp.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseFOptsUp(t *testing.T) {
	type args struct {
		payload []byte
	}
	tests := []struct {
		name string
		args args
		want []commands.Fopter
	}{
		{"basic",
			args{payload: []byte{}},
			[]commands.Fopter{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseFOptsUp(tt.args.payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFOptsUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
