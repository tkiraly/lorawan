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
		{"basic-with fpending",
			args{
				devaddr:  []byte{0x48, 0x52, 0x1f, 0xbd},
				fcnt:     2,
				fopts:    []commands.Fopter{},
				adr:      true,
				ack:      true,
				fpending: true,
				foptslen: 0,
			},
			fHDRDown([]byte{0xbd, 0x1f, 0x52, 0x48, 0xb0, 0x02, 0x00}),
		},
		{"basic-with devstatusreq",
			args{
				devaddr:  []byte{0x48, 0x52, 0x1f, 0xbd},
				fcnt:     2,
				fopts:    []commands.Fopter{devstatus.NewReq()},
				adr:      true,
				ack:      true,
				fpending: false,
				foptslen: 1,
			},
			fHDRDown([]byte{0xbd, 0x1f, 0x52, 0x48, 0xa1, 0x02, 0x00, 0x06}),
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
		{"basic with devstatusreq",
			fHDRDown([]byte{0xbd, 0x1f, 0x52, 0x48, 0xa1, 0x02, 0x00, 0x06}),
			"DevAddr: 48521FBD; FCnt: 2; ADR: true; FPending: false; ACK: true; FOptsLen: 1, FOpts: [DevStatusReq!]",
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
		{"basic-issues warning",
			args{payload: []byte{0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00}},
			[]commands.Fopter{},
		},
		{"linkcheckans",
			args{payload: []byte{0x02, 0x14, 0x02}},
			[]commands.Fopter{linkcheck.NewAns(20, 2)},
		},
		{"linkadrreq",
			args{payload: []byte{0x03, 0x24, 0x00, 0xff, 0x00}},
			[]commands.Fopter{linkadr.NewReq(2, 4, 0x00ff, 0, 0)},
		},
		{"dutycyclereq",
			args{payload: []byte{0x04, 0x05}},
			[]commands.Fopter{dutycycle.NewReq(5)},
		},
		{"rxparamsetupreq",
			args{payload: []byte{0x05, 0x35, 0x18, 0x4F, 0x84}},
			[]commands.Fopter{rxparamsetup.NewReq(3, 5, 8671000)},
		},
		{"devstatusreq",
			args{payload: []byte{0x06}},
			[]commands.Fopter{devstatus.NewReq()},
		},
		{"newchannelreq",
			args{payload: []byte{0x07, 0x05, 0x18, 0x4F, 0x84, 0x05}},
			[]commands.Fopter{newchannel.NewReq(5, 8671000, 0, 5)},
		},
		{"rxtimingsetupreq",
			args{payload: []byte{0x08, 0x05}},
			[]commands.Fopter{rxtimingsetup.NewReq(5)},
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
