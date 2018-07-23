package fhdr

import (
	"reflect"
	"testing"

	"github.com/tkiraly/lorawan/commands"
)

func Test_fHDRDown_DevAddr(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRDown
		want []byte
	}{
		{"basic",
			fHDRDown([]byte{0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00}),
			[]byte{0x48, 0x52, 0x1f, 0xbd},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.DevAddr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fHDRDown.DevAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRDown_FCnt(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRDown
		want uint16
	}{
		{"basic",
			fHDRDown([]byte{0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00}),
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.FCnt(); got != tt.want {
				t.Errorf("fHDRDown.FCnt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRDown_FOpts(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRDown
		want []commands.Fopter
	}{
		{"basic",
			fHDRDown([]byte{0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00}),
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.FOpts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fHDRDown.FOpts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRDown_FPending(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRDown
		want bool
	}{
		{"basic",
			fHDRDown([]byte{0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00}),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.FPending(); got != tt.want {
				t.Errorf("fHDRDown.FPending() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRDown_ADR(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRDown
		want bool
	}{
		{"basic",
			fHDRDown([]byte{0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.ADR(); got != tt.want {
				t.Errorf("fHDRDown.ADR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRDown_ACK(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRDown
		want bool
	}{
		{"basic",
			fHDRDown([]byte{0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.ACK(); got != tt.want {
				t.Errorf("fHDRDown.ACK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRDown_FOptsLen(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRDown
		want uint8
	}{
		{"basic",
			fHDRDown([]byte{0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00}),
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.FOptsLen(); got != tt.want {
				t.Errorf("fHDRDown.FOptsLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDown(t *testing.T) {
	type args struct {
		devaddr  []byte
		fcnt     uint16
		fopts    []commands.Fopter
		adr      bool
		ack      bool
		fpending bool
		foptslen uint8
	}
	tests := []struct {
		name string
		args args
		want FHDRDown
	}{
		{"basic",
			args{
				devaddr:  []byte{0x48, 0x52, 0x1f, 0xbd},
				fcnt:     2,
				fopts:    []commands.Fopter{},
				adr:      true,
				ack:      true,
				fpending: false,
				foptslen: 0,
			},
			fHDRDown([]byte{0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDown(tt.args.devaddr, tt.args.fcnt, tt.args.fopts, tt.args.adr, tt.args.ack, tt.args.fpending, tt.args.foptslen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRDown_String(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRDown
		want string
	}{
		{"basic",
			fHDRDown([]byte{0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00}),
			"DevAddr: 48521FBD; FCnt: 2; ADR: true; FPending: false; ACK: true; FOptsLen: 0, FOpts: []",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.String(); got != tt.want {
				t.Errorf("fHDRDown.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDown(t *testing.T) {
	type args struct {
		bb []byte
	}
	tests := []struct {
		name string
		args args
		want FHDRDown
	}{
		{"basic",
			args{bb: []byte{0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00}},
			fHDRDown([]byte{0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDown(tt.args.bb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fHDRDown_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		f    fHDRDown
		want []byte
	}{
		{"basic",
			fHDRDown([]byte{0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00}),
			[]byte{0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fHDRDown.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseFOptsDown(t *testing.T) {
	type args struct {
		payload []byte
	}
	tests := []struct {
		name string
		args args
		want []commands.Fopter
	}{
		{"basic",
			args{payload: []byte{0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00}},
			[]commands.Fopter{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseFOptsDown(tt.args.payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFOptsDown() = %v, want %v", got, tt.want)
			}
		})
	}
}
