package dlchannel

import (
	"reflect"
	"testing"

	"github.com/tkiraly/lorawan/commands"
)

func Test_dlChannelReq_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    dlChannelReq
		want []byte
	}{
		{"basic",
			dlChannelReq([]byte{commands.DlChannelReqCommand, 0x03, 0xE8, 0x56, 0x84}),
			[]byte{0x0A, 0x03, 0xE8, 0x56, 0x84},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dlChannelReq.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dlChannelReq_Len(t *testing.T) {
	tests := []struct {
		name string
		c    dlChannelReq
		want uint8
	}{
		{"basic",
			dlChannelReq([]byte{commands.DlChannelReqCommand, 0x03, 0xE8, 0x56, 0x84}),
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("dlChannelReq.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dlChannelReq_String(t *testing.T) {
	tests := []struct {
		name string
		c    dlChannelReq
		want string
	}{
		{"basic",
			dlChannelReq([]byte{commands.DlChannelReqCommand, 0x03, 0xE8, 0x56, 0x84}),
			"DlChannelReq! ChIndex: 3, Freq: 8673000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("dlChannelReq.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{0x0A, 0x03, 0xE8, 0x56, 0x84}},
			dlChannelReq([]byte{commands.DlChannelReqCommand, 0x03, 0xE8, 0x56, 0x84}),
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
		chindex uint8
		freq    uint32
	}
	tests := []struct {
		name string
		args args
		want DlChannelReq
	}{
		{"basic",
			args{chindex: 3,
				freq: 8673000,
			},
			dlChannelReq([]byte{commands.DlChannelReqCommand, 0x03, 0xE8, 0x56, 0x84}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReq(tt.args.chindex, tt.args.freq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dlChannelReq_ChIndex(t *testing.T) {
	tests := []struct {
		name string
		c    dlChannelReq
		want uint8
	}{
		{"basic",
			dlChannelReq([]byte{commands.DlChannelReqCommand, 0x03, 0xE8, 0x56, 0x84}),
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ChIndex(); got != tt.want {
				t.Errorf("dlChannelReq.ChIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dlChannelReq_Freq(t *testing.T) {
	tests := []struct {
		name string
		c    dlChannelReq
		want uint32
	}{
		{"basic",
			dlChannelReq([]byte{commands.DlChannelReqCommand, 0x03, 0xE8, 0x56, 0x84}),
			8673000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Freq(); got != tt.want {
				t.Errorf("dlChannelReq.Freq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dlChannelAns_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    dlChannelAns
		want []byte
	}{
		{"basic",
			dlChannelAns([]byte{commands.DlChannelReqCommand, 0x03}),
			[]byte{0x0A, 0x03},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dlChannelAns.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dlChannelAns_Len(t *testing.T) {
	tests := []struct {
		name string
		c    dlChannelAns
		want uint8
	}{
		{"basic",
			dlChannelAns([]byte{commands.DlChannelReqCommand, 0x03}),
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("dlChannelAns.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dlChannelAns_String(t *testing.T) {
	tests := []struct {
		name string
		c    dlChannelAns
		want string
	}{
		{"basic",
			dlChannelAns([]byte{commands.DlChannelReqCommand, 0x03}),
			"DlChannelAns! UplinkFrequencyExists: true, ChannelFrequencyOk: true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("dlChannelAns.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{0x0A, 0x03}},
			dlChannelAns([]byte{commands.DlChannelReqCommand, 0x03}),
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

func Test_dlChannelAns_ChannelFrequencyOk(t *testing.T) {
	tests := []struct {
		name string
		c    dlChannelAns
		want bool
	}{
		{"basic",
			dlChannelAns([]byte{commands.DlChannelReqCommand, 0x03}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ChannelFrequencyOk(); got != tt.want {
				t.Errorf("dlChannelAns.ChannelFrequencyOk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dlChannelAns_UplinkFrequencyExists(t *testing.T) {
	tests := []struct {
		name string
		c    dlChannelAns
		want bool
	}{
		{"basic",
			dlChannelAns([]byte{commands.DlChannelReqCommand, 0x03}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.UplinkFrequencyExists(); got != tt.want {
				t.Errorf("dlChannelAns.UplinkFrequencyExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAns(t *testing.T) {
	type args struct {
		uplinkfrequencyexists bool
		channelfrequencyok    bool
	}
	tests := []struct {
		name string
		args args
		want DlChannelAns
	}{
		{"basic",
			args{
				uplinkfrequencyexists: true,
				channelfrequencyok:    true,
			},
			dlChannelAns([]byte{commands.DlChannelReqCommand, 0x03}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAns(tt.args.uplinkfrequencyexists, tt.args.channelfrequencyok); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAns() = %v, want %v", got, tt.want)
			}
		})
	}
}
