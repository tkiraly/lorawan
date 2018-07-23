package util

import (
	"reflect"
	"testing"
)

func TestBytereverse(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"basic",
			args{
				s: []byte{0x04, 0x08},
			},
			[]byte{0x08, 0x04},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytereverse(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytereverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
