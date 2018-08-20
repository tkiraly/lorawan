package txparamsetup

import (
	"reflect"
	"testing"

	"github.com/tkiraly/lorawan/commands"
)

func Test_txParamSetupReq_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    txParamSetupReq
		want []byte
	}{
		{"basic",
			txParamSetupReq([]byte{commands.TxParamSetupReqCommand, 0x29}),
			[]byte{0x09, 0x29},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("txParamSetupReq.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_txParamSetupReq_Len(t *testing.T) {
	tests := []struct {
		name string
		c    txParamSetupReq
		want uint8
	}{
		{"basic",
			txParamSetupReq([]byte{commands.TxParamSetupReqCommand, 0x29}),
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("txParamSetupReq.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_txParamSetupReq_String(t *testing.T) {
	tests := []struct {
		name string
		c    txParamSetupReq
		want string
	}{
		{"basic",
			txParamSetupReq([]byte{commands.TxParamSetupReqCommand, 0x29}),
			"TxParamSetupReq! DownlinkDwellTime: 1, UplinkDwellTime: 0, MaxEIRP: 9;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("txParamSetupReq.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{0x09, 0x29}},
			txParamSetupReq([]byte{commands.TxParamSetupReqCommand, 0x29}),
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
		downlinkdwelltime uint8
		uplinkdwelltime   uint8
		maxeirp           uint8
	}
	tests := []struct {
		name string
		args args
		want TxParamSetupReq
	}{
		{"basic",
			args{
				downlinkdwelltime: 1,
				uplinkdwelltime:   0,
				maxeirp:           9,
			},
			txParamSetupReq([]byte{commands.TxParamSetupReqCommand, 0x29}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReq(tt.args.downlinkdwelltime, tt.args.uplinkdwelltime, tt.args.maxeirp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_txParamSetupReq_DownlinkDwellTime(t *testing.T) {
	tests := []struct {
		name string
		c    txParamSetupReq
		want uint8
	}{
		{"basic",
			txParamSetupReq([]byte{commands.TxParamSetupReqCommand, 0x29}),
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.DownlinkDwellTime(); got != tt.want {
				t.Errorf("txParamSetupReq.DownlinkDwellTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_txParamSetupReq_UplinkDwellTime(t *testing.T) {
	tests := []struct {
		name string
		c    txParamSetupReq
		want uint8
	}{
		{"basic",
			txParamSetupReq([]byte{commands.TxParamSetupReqCommand, 0x29}),
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.UplinkDwellTime(); got != tt.want {
				t.Errorf("txParamSetupReq.UplinkDwellTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_txParamSetupReq_MaxEIRP(t *testing.T) {
	tests := []struct {
		name string
		c    txParamSetupReq
		want uint8
	}{
		{"basic",
			txParamSetupReq([]byte{commands.TxParamSetupReqCommand, 0x29}),
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.MaxEIRP(); got != tt.want {
				t.Errorf("txParamSetupReq.MaxEIRP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_txParamSetupAns_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    txParamSetupAns
		want []byte
	}{
		{"basic",
			txParamSetupAns([]byte{commands.TxParamSetupAnsCommand}),
			[]byte{0x09},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("txParamSetupAns.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_txParamSetupAns_Len(t *testing.T) {
	tests := []struct {
		name string
		c    txParamSetupAns
		want uint8
	}{
		{"basic",
			txParamSetupAns([]byte{commands.TxParamSetupAnsCommand}),
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("txParamSetupAns.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_txParamSetupAns_String(t *testing.T) {
	tests := []struct {
		name string
		c    txParamSetupAns
		want string
	}{
		{"basic",
			txParamSetupAns([]byte{commands.TxParamSetupAnsCommand}),
			"TxParamSetupAns!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("txParamSetupAns.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{0x09}},
			txParamSetupAns([]byte{commands.TxParamSetupAnsCommand}),
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
	tests := []struct {
		name string
		want TxParamSetupAns
	}{
		{"basic",
			txParamSetupAns([]byte{commands.TxParamSetupAnsCommand}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAns() = %v, want %v", got, tt.want)
			}
		})
	}
}
