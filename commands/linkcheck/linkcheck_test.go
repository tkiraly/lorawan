package linkcheck

import (
	"reflect"
	"testing"

	"github.com/tkiraly/lorawan/commands"
)

func Test_linkCheckReq_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    linkCheckReq
		want []byte
	}{
		{"basic",
			linkCheckReq([]byte{0x02}),
			[]byte{0x02},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("linkCheckReq.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkCheckReq_Len(t *testing.T) {
	tests := []struct {
		name string
		c    linkCheckReq
		want uint8
	}{
		{"basic",
			linkCheckReq([]byte{0x02}),
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("linkCheckReq.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkCheckReq_String(t *testing.T) {
	tests := []struct {
		name string
		c    linkCheckReq
		want string
	}{
		{"basic",
			linkCheckReq([]byte{0x02}),
			"LinkCheckReq!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("linkCheckReq.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{0x02}},
			linkCheckReq([]byte{0x02}),
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
	tests := []struct {
		name string
		want LinkCheckReq
	}{
		{"basic",
			linkCheckReq([]byte{0x02}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkCheckAns_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    linkCheckAns
		want []byte
	}{
		{"basic",
			linkCheckAns([]byte{0x02, 0x20, 0x04}),
			[]byte{0x02, 0x20, 0x04},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("linkCheckAns.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkCheckAns_Len(t *testing.T) {
	tests := []struct {
		name string
		c    linkCheckAns
		want uint8
	}{
		{"basic",
			linkCheckAns([]byte{0x02, 0x20, 0x04}),
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("linkCheckAns.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkCheckAns_Margin(t *testing.T) {
	tests := []struct {
		name string
		c    linkCheckAns
		want uint8
	}{
		{"basic",
			linkCheckAns([]byte{0x02, 0x20, 0x04}),
			0x20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Margin(); got != tt.want {
				t.Errorf("linkCheckAns.Margin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkCheckAns_GwCnt(t *testing.T) {
	tests := []struct {
		name string
		c    linkCheckAns
		want uint8
	}{
		{"basic",
			linkCheckAns([]byte{0x02, 0x20, 0x04}),
			0x04,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GwCnt(); got != tt.want {
				t.Errorf("linkCheckAns.GwCnt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkCheckAns_String(t *testing.T) {
	tests := []struct {
		name string
		c    linkCheckAns
		want string
	}{
		{"basic",
			linkCheckAns([]byte{0x02, 0x20, 0x04}),
			"LinkCheckAns! Margin: 32, GwCnt: 4;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("linkCheckAns.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{0x02, 0x20, 0x04}},
			linkCheckAns([]byte{0x02, 0x20, 0x04}),
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
		margin uint8
		gwcnt  uint8
	}
	tests := []struct {
		name string
		args args
		want LinkCheckAns
	}{
		{"basic",
			args{margin: 0x20, gwcnt: 0x04},
			linkCheckAns([]byte{0x02, 0x20, 0x04}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAns(tt.args.margin, tt.args.gwcnt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAns() = %v, want %v", got, tt.want)
			}
		})
	}
}
