package newchannel

import (
	"reflect"
	"testing"

	"github.com/tkiraly/lorawan/commands"
)

func Test_newChannelReq_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    newChannelReq
		want []byte
	}{
		{"basic",
			newChannelReq([]byte{commands.NewChannelReqCommand, 0x02, 0x18, 0x4F, 0x84, 0x70}),
			[]byte{0x07, 0x02, 0x18, 0x4F, 0x84, 0x70},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newChannelReq.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newChannelReq_Len(t *testing.T) {
	tests := []struct {
		name string
		c    newChannelReq
		want uint8
	}{
		{"basic",
			newChannelReq([]byte{commands.NewChannelReqCommand, 0x02, 0x18, 0x4F, 0x84, 0x70}),
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("newChannelReq.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newChannelReq_String(t *testing.T) {
	tests := []struct {
		name string
		c    newChannelReq
		want string
	}{
		{"basic",
			newChannelReq([]byte{commands.NewChannelReqCommand, 0x02, 0x18, 0x4F, 0x84, 0x70}),
			"NewChannelReq! ChIndex: 2, Freq: 8671000, MaxDR: 7, MinDR: 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("newChannelReq.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{commands.NewChannelReqCommand, 0x02, 0x18, 0x4F, 0x84, 0x70}},
			newChannelReq([]byte{commands.NewChannelReqCommand, 0x02, 0x18, 0x4F, 0x84, 0x70}),
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
		maxdr   uint8
		mindr   uint8
	}
	tests := []struct {
		name string
		args args
		want NewChannelReq
	}{
		{"basic",
			args{chindex: 0x02, freq: 8671000, maxdr: 7, mindr: 0},
			newChannelReq([]byte{commands.NewChannelReqCommand, 0x02, 0x18, 0x4F, 0x84, 0x70}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReq(tt.args.chindex, tt.args.freq, tt.args.maxdr, tt.args.mindr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newChannelReq_ChIndex(t *testing.T) {
	tests := []struct {
		name string
		c    newChannelReq
		want uint8
	}{
		{"basic",
			newChannelReq([]byte{commands.NewChannelReqCommand, 0x02, 0x18, 0x4F, 0x84, 0x70}),
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ChIndex(); got != tt.want {
				t.Errorf("newChannelReq.ChIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newChannelReq_Freq(t *testing.T) {
	tests := []struct {
		name string
		c    newChannelReq
		want uint32
	}{
		{"basic",
			newChannelReq([]byte{commands.NewChannelReqCommand, 0x02, 0x18, 0x4F, 0x84, 0x70}),
			8671000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Freq(); got != tt.want {
				t.Errorf("newChannelReq.Freq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newChannelReq_MaxDR(t *testing.T) {
	tests := []struct {
		name string
		c    newChannelReq
		want uint8
	}{
		{"basic",
			newChannelReq([]byte{commands.NewChannelReqCommand, 0x02, 0x18, 0x4F, 0x84, 0x70}),
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.MaxDR(); got != tt.want {
				t.Errorf("newChannelReq.MaxDR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newChannelReq_MinDR(t *testing.T) {
	tests := []struct {
		name string
		c    newChannelReq
		want uint8
	}{
		{"basic",
			newChannelReq([]byte{commands.NewChannelReqCommand, 0x02, 0x18, 0x4F, 0x84, 0x70}),
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.MinDR(); got != tt.want {
				t.Errorf("newChannelReq.MinDR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newChannelAns_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    newChannelAns
		want []byte
	}{
		{"basic",
			newChannelAns([]byte{commands.NewChannelAnsCommand, 0x02}),
			[]byte{0x07, 0x02},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newChannelAns.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newChannelAns_Len(t *testing.T) {
	tests := []struct {
		name string
		c    newChannelAns
		want uint8
	}{
		{"basic",
			newChannelAns([]byte{commands.NewChannelAnsCommand, 0x02}),
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("newChannelAns.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newChannelAns_String(t *testing.T) {
	tests := []struct {
		name string
		c    newChannelAns
		want string
	}{
		{"basic",
			newChannelAns([]byte{commands.NewChannelAnsCommand, 0x02}),
			"NewChannelAns! DataRateRangeOK: true, ChannelFrequencyOK: false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("newChannelAns.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{0x07, 0x02}},
			newChannelAns([]byte{commands.NewChannelAnsCommand, 0x02}),
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
		dataraterangeok    bool
		channelfrequencyok bool
	}
	tests := []struct {
		name string
		args args
		want NewChannelAns
	}{
		{"basic",
			args{dataraterangeok: true, channelfrequencyok: false},
			newChannelAns([]byte{commands.NewChannelAnsCommand, 0x02}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAns(tt.args.dataraterangeok, tt.args.channelfrequencyok); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newChannelAns_DataRateRangeOK(t *testing.T) {
	tests := []struct {
		name string
		c    newChannelAns
		want bool
	}{
		{"basic",
			newChannelAns([]byte{commands.NewChannelAnsCommand, 0x02}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.DataRateRangeOK(); got != tt.want {
				t.Errorf("newChannelAns.DataRateRangeOK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newChannelAns_ChannelFrequencyOK(t *testing.T) {
	tests := []struct {
		name string
		c    newChannelAns
		want bool
	}{
		{"basic",
			newChannelAns([]byte{commands.NewChannelAnsCommand, 0x02}),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ChannelFrequencyOK(); got != tt.want {
				t.Errorf("newChannelAns.ChannelFrequencyOK() = %v, want %v", got, tt.want)
			}
		})
	}
}
