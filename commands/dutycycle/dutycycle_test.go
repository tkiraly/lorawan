package dutycycle

import (
	"reflect"
	"testing"

	"github.com/tkiraly/lorawan/commands"
)

func Test_dutyCycleReq_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    dutyCycleReq
		want []byte
	}{
		{"basic",
			dutyCycleReq([]byte{commands.DutyCycleReqCommand, 0x0f}),
			[]byte{0x04, 0x0f},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dutyCycleReq.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dutyCycleReq_Len(t *testing.T) {
	tests := []struct {
		name string
		c    dutyCycleReq
		want uint8
	}{
		{"basic",
			dutyCycleReq([]byte{commands.DutyCycleReqCommand, 0x0f}),
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("dutyCycleReq.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dutyCycleReq_String(t *testing.T) {
	tests := []struct {
		name string
		c    dutyCycleReq
		want string
	}{
		{"basic",
			dutyCycleReq([]byte{commands.DutyCycleReqCommand, 0x0f}),
			"DutyCycleReq! MaxDCycle: 15",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("dutyCycleReq.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{0x04, 0x0f}},
			dutyCycleReq([]byte{commands.DutyCycleReqCommand, 0x0f}),
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
		maxdcycle uint8
	}
	tests := []struct {
		name string
		args args
		want DutyCycleReq
	}{
		{"basic",
			args{maxdcycle: 15},
			dutyCycleReq([]byte{commands.DutyCycleReqCommand, 0x0f}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReq(tt.args.maxdcycle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dutyCycleReq_MaxDCycle(t *testing.T) {
	tests := []struct {
		name string
		c    dutyCycleReq
		want uint8
	}{
		{"basic",
			dutyCycleReq([]byte{commands.DutyCycleReqCommand, 0x0f}),
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.MaxDCycle(); got != tt.want {
				t.Errorf("dutyCycleReq.MaxDCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dutyCycleAns_ByteArray(t *testing.T) {
	tests := []struct {
		name string
		c    dutyCycleAns
		want []byte
	}{
		{"basic",
			dutyCycleAns([]byte{commands.DutyCycleAnsCommand}),
			[]byte{0x04},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ByteArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dutyCycleAns.ByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dutyCycleAns_Len(t *testing.T) {
	tests := []struct {
		name string
		c    dutyCycleAns
		want uint8
	}{
		{"basic",
			dutyCycleAns([]byte{commands.DutyCycleAnsCommand}),
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Len(); got != tt.want {
				t.Errorf("dutyCycleAns.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dutyCycleAns_String(t *testing.T) {
	tests := []struct {
		name string
		c    dutyCycleAns
		want string
	}{
		{"basic",
			dutyCycleAns([]byte{commands.DutyCycleAnsCommand}),
			"DutyCycleAns!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("dutyCycleAns.String() = %v, want %v", got, tt.want)
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
			args{bb: []byte{0x04}},
			dutyCycleAns([]byte{commands.DutyCycleAnsCommand}),
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
		want DutyCycleAns
	}{
		{"basic",
			dutyCycleAns([]byte{commands.DutyCycleAnsCommand}),
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
