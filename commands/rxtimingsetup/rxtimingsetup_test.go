package rxtimingsetup

import (
	"reflect"
	"testing"

	"github.com/tkiraly/lorawan/commands"
)

func Test_rxTimingSetupReq_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    rxTimingSetupReq
		want []byte
	}{
		{"basic",
			rxTimingSetupReq([]byte{commands.RxTimingSetupReqCommand, 0x0f}),
			[]byte{0x08, 0x0f},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rxTimingSetupReq.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rxTimingSetupReq_Len(t *testing.T) {
	tests := []struct {
		name string
		c    rxTimingSetupReq
		want uint8
	}{
		{"basic",
			rxTimingSetupReq([]byte{commands.RxTimingSetupReqCommand, 0x0f}),
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("rxTimingSetupReq.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rxTimingSetupReq_String(t *testing.T) {
	tests := []struct {
		name string
		c    rxTimingSetupReq
		want string
	}{
		{"basic",
			rxTimingSetupReq([]byte{commands.RxTimingSetupReqCommand, 0x0f}),
			"RxTimingSetupReq! Del: 15",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("rxTimingSetupReq.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{0x08, 0x0f}},
			rxTimingSetupReq([]byte{commands.RxTimingSetupReqCommand, 0x0f}),
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
		del uint8
	}
	tests := []struct {
		name string
		args args
		want RxTimingSetupReq
	}{
		{"basic",
			args{del: 15},
			rxTimingSetupReq([]byte{commands.RxTimingSetupReqCommand, 0x0f}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReq(tt.args.del); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rxTimingSetupReq_Del(t *testing.T) {
	tests := []struct {
		name string
		c    rxTimingSetupReq
		want uint8
	}{
		{"basic",
			rxTimingSetupReq([]byte{commands.RxTimingSetupReqCommand, 0x0f}),
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Del(); got != tt.want {
				t.Errorf("rxTimingSetupReq.Del() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rxTimingSetupAns_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    rxTimingSetupAns
		want []byte
	}{
		{"basic",
			rxTimingSetupAns([]byte{commands.RxTimingSetupAnsCommand}),
			[]byte{0x08},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rxTimingSetupAns.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rxTimingSetupAns_Len(t *testing.T) {
	tests := []struct {
		name string
		c    rxTimingSetupAns
		want uint8
	}{
		{"basic",
			rxTimingSetupAns([]byte{commands.RxTimingSetupAnsCommand}),
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("rxTimingSetupAns.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rxTimingSetupAns_String(t *testing.T) {
	tests := []struct {
		name string
		c    rxTimingSetupAns
		want string
	}{
		{"basic",
			rxTimingSetupAns([]byte{commands.RxTimingSetupAnsCommand}),
			"RxTimingSetupAns!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("rxTimingSetupAns.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{0x08}},
			rxTimingSetupAns([]byte{commands.RxTimingSetupAnsCommand}),
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
		want RxTimingSetupAns
	}{
		{"basic",
			rxTimingSetupAns([]byte{commands.RxTimingSetupAnsCommand}),
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
