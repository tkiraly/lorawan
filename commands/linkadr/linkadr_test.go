package linkadr

import (
	"reflect"
	"testing"

	"github.com/tkiraly/lorawan/commands"
)

func Test_linkADRReq_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    linkADRReq
		want []byte
	}{
		{"basic",
			linkADRReq([]byte{commands.LinkADRReqCommand, 0x22, 0xff, 0x00, 0x00}),
			[]byte{commands.LinkADRReqCommand, 0x22, 0xff, 0x00, 0x00},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("linkADRReq.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkADRReq_Len(t *testing.T) {
	tests := []struct {
		name string
		c    linkADRReq
		want uint8
	}{
		{"basic",
			linkADRReq([]byte{commands.LinkADRReqCommand, 0x22, 0xff, 0x00, 0x00}),
			Reqlen,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("linkADRReq.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkADRReq_String(t *testing.T) {
	tests := []struct {
		name string
		c    linkADRReq
		want string
	}{
		{"basic",
			linkADRReq([]byte{commands.LinkADRReqCommand, 0x22, 0xff, 0x00, 0x00}),
			"LinkADRReq! Datarate: 2, TxPower: 2, Chmask: ff00, ChmaskCtrl: 0, Nbtrans: 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("linkADRReq.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{commands.LinkADRReqCommand, 0x22, 0xff, 0x00, 0x00}},
			linkADRReq([]byte{commands.LinkADRReqCommand, 0x22, 0xff, 0x00, 0x00}),
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
		datarate   uint8
		txpower    uint8
		chmask     uint16
		chmaskcntl uint8
		nbtrans    uint8
	}
	tests := []struct {
		name string
		args args
		want LinkADRReq
	}{
		{"basic",
			args{
				datarate:   2,
				txpower:    2,
				chmask:     0xff00,
				chmaskcntl: 0,
				nbtrans:    0,
			},
			linkADRReq([]byte{commands.LinkADRReqCommand, 0x22, 0xff, 0x00, 0x00}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReq(tt.args.datarate, tt.args.txpower, tt.args.chmask, tt.args.chmaskcntl, tt.args.nbtrans); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkADRReq_DataRate(t *testing.T) {
	tests := []struct {
		name string
		c    linkADRReq
		want uint8
	}{
		{"basic",
			linkADRReq([]byte{commands.LinkADRReqCommand, 0x22, 0xff, 0x00, 0x00}),
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.DataRate(); got != tt.want {
				t.Errorf("linkADRReq.DataRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkADRReq_TxPower(t *testing.T) {
	tests := []struct {
		name string
		c    linkADRReq
		want uint8
	}{
		{"basic",
			linkADRReq([]byte{commands.LinkADRReqCommand, 0x22, 0xff, 0x00, 0x00}),
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.TxPower(); got != tt.want {
				t.Errorf("linkADRReq.TxPower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkADRReq_ChMask(t *testing.T) {
	tests := []struct {
		name string
		c    linkADRReq
		want uint16
	}{
		{"basic",
			linkADRReq([]byte{commands.LinkADRReqCommand, 0x22, 0xff, 0x00, 0x00}),
			0xff00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ChMask(); got != tt.want {
				t.Errorf("linkADRReq.ChMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkADRReq_Redundancy_ChMaskCntl(t *testing.T) {
	tests := []struct {
		name string
		c    linkADRReq
		want uint8
	}{
		{"basic",
			linkADRReq([]byte{commands.LinkADRReqCommand, 0x22, 0xff, 0x00, 0x00}),
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Redundancy_ChMaskCntl(); got != tt.want {
				t.Errorf("linkADRReq.Redundancy_ChMaskCntl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkADRReq_Redundancy_NbTrans(t *testing.T) {
	tests := []struct {
		name string
		c    linkADRReq
		want uint8
	}{
		{"basic",
			linkADRReq([]byte{commands.LinkADRReqCommand, 0x22, 0xff, 0x00, 0x00}),
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Redundancy_NbTrans(); got != tt.want {
				t.Errorf("linkADRReq.Redundancy_NbTrans() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkADRAns_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    linkADRAns
		want []byte
	}{
		{"basic",
			linkADRAns([]byte{commands.LinkADRAnsCommand, 0x07}),
			[]byte{commands.LinkADRAnsCommand, 0x07},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("linkADRAns.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkADRAns_Len(t *testing.T) {
	tests := []struct {
		name string
		c    linkADRAns
		want uint8
	}{
		{"basic",
			linkADRAns([]byte{commands.LinkADRAnsCommand, 0x07}),
			Anslen,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("linkADRAns.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkADRAns_PowerACK(t *testing.T) {
	tests := []struct {
		name string
		c    linkADRAns
		want bool
	}{
		{"basic",
			linkADRAns([]byte{commands.LinkADRAnsCommand, 0x07}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.PowerACK(); got != tt.want {
				t.Errorf("linkADRAns.PowerACK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkADRAns_DatarateACK(t *testing.T) {
	tests := []struct {
		name string
		c    linkADRAns
		want bool
	}{
		{"basic",
			linkADRAns([]byte{commands.LinkADRAnsCommand, 0x07}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.DatarateACK(); got != tt.want {
				t.Errorf("linkADRAns.DatarateACK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkADRAns_ChannelmaskACK(t *testing.T) {
	tests := []struct {
		name string
		c    linkADRAns
		want bool
	}{
		{"basic",
			linkADRAns([]byte{commands.LinkADRAnsCommand, 0x07}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ChannelmaskACK(); got != tt.want {
				t.Errorf("linkADRAns.ChannelmaskACK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkADRAns_String(t *testing.T) {
	tests := []struct {
		name string
		c    linkADRAns
		want string
	}{
		{"basic",
			linkADRAns([]byte{commands.LinkADRAnsCommand, 0x07}),
			"LinkADRAns! PowerACK: true, DatarateACK: true, ChannelmaskACK: true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("linkADRAns.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{commands.LinkADRAnsCommand, 0x07}},
			linkADRAns([]byte{commands.LinkADRAnsCommand, 0x07}),
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

func TestNewAns(t *testing.T) {
	type args struct {
		powerack       bool
		datarateack    bool
		channelmaskack bool
	}
	tests := []struct {
		name string
		args args
		want LinkADRAns
	}{
		{"basic",
			args{powerack: true, datarateack: true, channelmaskack: true},
			linkADRAns([]byte{commands.LinkADRAnsCommand, 0x07}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAns(tt.args.powerack, tt.args.datarateack, tt.args.channelmaskack); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAns() = %v, want %v", got, tt.want)
			}
		})
	}
}
