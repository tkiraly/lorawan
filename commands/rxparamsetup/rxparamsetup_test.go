package rxparamsetup

import (
	"reflect"
	"testing"

	"github.com/tkiraly/lorawan/commands"
)

func Test_rXParamSetupReq_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    rXParamSetupReq
		want []byte
	}{
		{"basic",
			rXParamSetupReq([]byte{commands.RXParamSetupReqCommand, 0x00, 0x18, 0x4F, 0x84}),
			[]byte{commands.RXParamSetupReqCommand, 0x00, 0x18, 0x4F, 0x84},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rXParamSetupReq.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rXParamSetupReq_Len(t *testing.T) {
	tests := []struct {
		name string
		c    rXParamSetupReq
		want uint8
	}{
		{"basic",
			rXParamSetupReq([]byte{commands.RXParamSetupReqCommand, 0x00, 0x18, 0x4F, 0x84}),
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("rXParamSetupReq.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rXParamSetupReq_String(t *testing.T) {
	tests := []struct {
		name string
		c    rXParamSetupReq
		want string
	}{
		{"basic",
			rXParamSetupReq([]byte{commands.RXParamSetupReqCommand, 0x00, 0x18, 0x4F, 0x84}),
			"RXParamSetupReq! RX1DRoffset: 0, RX2DataRate: 0, Frequency: 8671000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("rXParamSetupReq.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rXParamSetupReq_RX1DRoffset(t *testing.T) {
	tests := []struct {
		name string
		c    rXParamSetupReq
		want uint8
	}{
		{"basic",
			rXParamSetupReq([]byte{commands.RXParamSetupReqCommand, 0x00, 0x18, 0x4F, 0x84}),
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.RX1DRoffset(); got != tt.want {
				t.Errorf("rXParamSetupReq.RX1DRoffset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rXParamSetupReq_RX2DataRate(t *testing.T) {
	tests := []struct {
		name string
		c    rXParamSetupReq
		want uint8
	}{
		{"basic",
			rXParamSetupReq([]byte{commands.RXParamSetupReqCommand, 0x00, 0x18, 0x4F, 0x84}),
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.RX2DataRate(); got != tt.want {
				t.Errorf("rXParamSetupReq.RX2DataRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rXParamSetupReq_Frequency(t *testing.T) {
	tests := []struct {
		name string
		c    rXParamSetupReq
		want uint32
	}{
		{"basic",
			rXParamSetupReq([]byte{commands.RXParamSetupReqCommand, 0x00, 0x18, 0x4F, 0x84}),
			8671000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Frequency(); got != tt.want {
				t.Errorf("rXParamSetupReq.Frequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseReq(t *testing.T) {
	type args struct {
		bb []byte
	}
	tests := []struct {
		name string
		args args
		want commands.Fopter
	}{
		{"basic",
			args{bb: []byte{commands.RXParamSetupReqCommand, 0x00, 0x18, 0x4F, 0x84}},
			rXParamSetupReq([]byte{commands.RXParamSetupReqCommand, 0x00, 0x18, 0x4F, 0x84}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseReq(tt.args.bb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewReq(t *testing.T) {
	type args struct {
		rx1droffset uint8
		rx2datarate uint8
		frequency   uint32
	}
	tests := []struct {
		name string
		args args
		want RXParamSetupReq
	}{
		{"basic",
			args{rx1droffset: 0, rx2datarate: 0, frequency: 8671000},
			rXParamSetupReq([]byte{commands.RXParamSetupReqCommand, 0x00, 0x18, 0x4F, 0x84}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReq(tt.args.rx1droffset, tt.args.rx2datarate, tt.args.frequency); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rXParamSetupAns_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    rXParamSetupAns
		want []byte
	}{
		{"basic",
			rXParamSetupAns([]byte{commands.RXParamSetupAnsCommand, 0x07}),
			[]byte{commands.RXParamSetupAnsCommand, 0x07},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rXParamSetupAns.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rXParamSetupAns_Len(t *testing.T) {
	tests := []struct {
		name string
		c    rXParamSetupAns
		want uint8
	}{
		{"basic",
			rXParamSetupAns([]byte{commands.RXParamSetupAnsCommand, 0x07}),
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("rXParamSetupAns.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rXParamSetupAns_String(t *testing.T) {
	tests := []struct {
		name string
		c    rXParamSetupAns
		want string
	}{
		{"basic",
			rXParamSetupAns([]byte{commands.RXParamSetupAnsCommand, 0x07}),
			"RXParamSetupAns! RX1DRoffsetACK: true, RX2DatarateACK: true, ChannelACK: true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("rXParamSetupAns.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAns(t *testing.T) {
	type args struct {
		bb []byte
	}
	tests := []struct {
		name string
		args args
		want commands.Fopter
	}{
		{"basic",
			args{bb: []byte{commands.RXParamSetupAnsCommand, 0x07}},
			rXParamSetupAns([]byte{commands.RXParamSetupAnsCommand, 0x07}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseAns(tt.args.bb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseAns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rXParamSetupAns_RX1DRoffsetACK(t *testing.T) {
	tests := []struct {
		name string
		c    rXParamSetupAns
		want bool
	}{
		{"basic",
			rXParamSetupAns([]byte{commands.RXParamSetupAnsCommand, 0x07}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.RX1DRoffsetACK(); got != tt.want {
				t.Errorf("rXParamSetupAns.RX1DRoffsetACK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rXParamSetupAns_RX2DatarateACK(t *testing.T) {
	tests := []struct {
		name string
		c    rXParamSetupAns
		want bool
	}{
		{"basic",
			rXParamSetupAns([]byte{commands.RXParamSetupAnsCommand, 0x07}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.RX2DatarateACK(); got != tt.want {
				t.Errorf("rXParamSetupAns.RX2DatarateACK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rXParamSetupAns_ChannelACK(t *testing.T) {
	tests := []struct {
		name string
		c    rXParamSetupAns
		want bool
	}{
		{"basic",
			rXParamSetupAns([]byte{commands.RXParamSetupAnsCommand, 0x07}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ChannelACK(); got != tt.want {
				t.Errorf("rXParamSetupAns.ChannelACK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAns(t *testing.T) {
	type args struct {
		rx1droffsetack bool
		rx2datarateack bool
		channelack     bool
	}
	tests := []struct {
		name string
		args args
		want RXParamSetupAns
	}{
		{"basic",
			args{rx1droffsetack: true, rx2datarateack: true, channelack: true},
			rXParamSetupAns([]byte{commands.RXParamSetupAnsCommand, 0x07}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAns(tt.args.rx1droffsetack, tt.args.rx2datarateack, tt.args.channelack); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAns() = %v, want %v", got, tt.want)
			}
		})
	}
}
