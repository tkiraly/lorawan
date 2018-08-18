package devstatus

import (
	"reflect"
	"testing"

	"github.com/tkiraly/lorawan/commands"
)

func Test_devStatusReq_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    devStatusReq
		want []byte
	}{
		{"basic",
			devStatusReq([]byte{commands.DevStatusReqCommand}),
			[]byte{0x06},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("devStatusReq.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_devStatusReq_Len(t *testing.T) {
	tests := []struct {
		name string
		c    devStatusReq
		want uint8
	}{
		{"basic",
			devStatusReq([]byte{commands.DevStatusReqCommand}),
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("devStatusReq.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_devStatusReq_String(t *testing.T) {
	tests := []struct {
		name string
		c    devStatusReq
		want string
	}{
		{"basic",
			devStatusReq([]byte{commands.DevStatusReqCommand}),
			"DevStatusReq!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("devStatusReq.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{0x06}},
			devStatusReq([]byte{commands.DevStatusReqCommand}),
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
		want DevStatusReq
	}{
		{"basic",
			devStatusReq([]byte{commands.DevStatusReqCommand}),
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

func Test_devStatusAns_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    devStatusAns
		want []byte
	}{
		{"basic",
			devStatusAns([]byte{commands.DevStatusAnsCommand, 0x56, 0x11}),
			[]byte{0x06, 0x56, 0x11},
		},
		{"basic2",
			devStatusAns([]byte{commands.DevStatusAnsCommand, 0x11, 0x11}),
			[]byte{0x06, 0x11, 0x11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("devStatusAns.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_devStatusAns_Len(t *testing.T) {
	tests := []struct {
		name string
		c    devStatusAns
		want uint8
	}{
		{"basic",
			devStatusAns([]byte{commands.DevStatusAnsCommand, 0x56, 0x11}),
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("devStatusAns.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_devStatusAns_Battery(t *testing.T) {
	tests := []struct {
		name string
		c    devStatusAns
		want uint8
	}{
		{"basic",
			devStatusAns([]byte{commands.DevStatusAnsCommand, 0x56, 0x11}),
			0x56,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Battery(); got != tt.want {
				t.Errorf("devStatusAns.Battery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_devStatusAns_Margin(t *testing.T) {
	tests := []struct {
		name string
		c    devStatusAns
		want uint8
	}{
		{"basic",
			devStatusAns([]byte{commands.DevStatusAnsCommand, 0x56, 0x11}),
			0x11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Margin(); got != tt.want {
				t.Errorf("devStatusAns.Margin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_devStatusAns_String(t *testing.T) {
	tests := []struct {
		name string
		c    devStatusAns
		want string
	}{
		{"basic",
			devStatusAns([]byte{commands.DevStatusAnsCommand, 0x56, 0x11}),
			"DevStatusAns! Battery: 86, Margin: 17",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("devStatusAns.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{0x06, 0x56, 0x11}},
			devStatusAns([]byte{commands.DevStatusAnsCommand, 0x56, 0x11}),
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
		battery uint8
		margin  uint8
	}
	tests := []struct {
		name string
		args args
		want DevStatusAns
	}{
		{"basic",
			args{battery: 0x56, margin: 0x11},
			devStatusAns([]byte{commands.DevStatusAnsCommand, 0x56, 0x11}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAns(tt.args.battery, tt.args.margin); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAns() = %v, want %v", got, tt.want)
			}
		})
	}
}
