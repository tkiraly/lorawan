package fhdr

import (
	"reflect"
	"testing"

	"github.com/tkiraly/lorawan/commands/devstatus"
	"github.com/tkiraly/lorawan/commands/dutycycle"
	"github.com/tkiraly/lorawan/commands/linkadr"
	"github.com/tkiraly/lorawan/commands/linkcheck"
	"github.com/tkiraly/lorawan/commands/newchannel"
	"github.com/tkiraly/lorawan/commands/rxparamsetup"
	"github.com/tkiraly/lorawan/commands/rxtimingsetup"

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
		{"basic with fopts",
			fHDRUp([]byte{0xbd, 0x1f, 0x52, 0x48, 0x81, 0x04, 0x00, 0x02}),
			[]commands.Fopter{linkcheck.NewReq()},
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
		{"basic with ack & adrackreq",

			args{
				devaddr:   []byte{0x48, 0x52, 0x1f, 0xbd},
				fcnt:      4,
				fopts:     []commands.Fopter{},
				adr:       true,
				ack:       true,
				adrackreq: true,
				foptslen:  0,
			},
			fHDRUp([]byte{0xbd, 0x1f, 0x52, 0x48, 0xe0, 0x04, 0x00}),
		},
		{"basic with ack & adrackreq",

			args{
				devaddr:   []byte{0x48, 0x52, 0x1f, 0xbd},
				fcnt:      4,
				fopts:     []commands.Fopter{linkcheck.NewReq()},
				adr:       true,
				ack:       true,
				adrackreq: true,
				foptslen:  1,
			},
			fHDRUp([]byte{0xbd, 0x1f, 0x52, 0x48, 0xe1, 0x04, 0x00, 0x02}),
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
		{"basic-with warning",
			args{payload: []byte{0x00}},
			[]commands.Fopter{},
		},
		{"linkcheck",
			args{payload: []byte{0x02}},
			[]commands.Fopter{linkcheck.NewReq()},
		},
		{"linkadr",
			args{payload: []byte{0x03, 0x06}},
			[]commands.Fopter{linkadr.NewAns(true, true, false)},
		},
		{"dutycycle",
			args{payload: []byte{0x04}},
			[]commands.Fopter{dutycycle.NewAns()},
		},
		{"rxparamsetup",
			args{payload: []byte{0x05, 0x02}},
			[]commands.Fopter{rxparamsetup.NewAns(false, true, false)},
		},
		{"devstatus",
			args{payload: []byte{0x06, 0x66, 0x02}},
			[]commands.Fopter{devstatus.NewAns(0x66, 0x02)},
		},
		{"newchannel",
			args{payload: []byte{0x07, 0x02}},
			[]commands.Fopter{newchannel.NewAns(true, false)},
		},
		{"rxtimingsetup",
			args{payload: []byte{0x08}},
			[]commands.Fopter{rxtimingsetup.NewAns()},
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
