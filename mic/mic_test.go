package mic

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	type args struct {
		payload []byte
		key     []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"joinreq",
			args{
				payload: []byte{
					0x00, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
					0x01, 0x54, 0x67, 0x1c, 0x00, 0x0b, 0xa3, 0x04,
					0x00, 0x4b, 0x33, 0x00, 0x00, 0x00, 0x00},
				key: []byte{
					0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
					0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
				},
			},
			[]byte{0xCD, 0xC8, 0xB4, 0x80},
			false,
		},
		{"joinans",
			args{
				/*payload: []byte{//encrypted
					0x20, 0xef, 0xef, 0xd5, 0xd7, 0xa3, 0xd0, 0x33,
					0x8a, 0xc5, 0xf6, 0x30, 0x66, 0x71, 0x19, 0xe2,
					0x08, 0x21, 0x58, 0x1e, 0xc6, 0xdd, 0xea, 0x95,
					0xc9, 0xff, 0xda, 0xef, 0x59, 0x00, 0x00, 0x00,
					0x00,
				},*/
				payload: []byte{ //unencrypted
					0x20, 0xa2, 0x12, 0x42, 0x24, 0x00, 0x00, 0xbd,
					0x1f, 0x52, 0x48, 0x00, 0x00, 0x18, 0x4f, 0x84,
					0xe8, 0x56, 0x84, 0xb8, 0x5e, 0x84, 0x88, 0x66,
					0x84, 0x58, 0x6e, 0x84, 0x00, 0x07, 0x62, 0x37,
					0x9b,
				},
				key: []byte{0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02},
			},
			[]byte{0x07, 0x62, 0x37, 0x9b},
			false,
		},
		{"dataupuncnf",
			args{
				payload: []byte{
					0x40, 0xbd, 0x1f, 0x52, 0x48, 0x80, 0x03, 0x00, 0x07, 0x2b, 0x00, 0x00, 0x00, 0x00,
				},
				key: []byte{0x7E, 0x49, 0x1B, 0x08, 0xF3, 0x09, 0xF0, 0x0B, 0xA7, 0xF2, 0xEE, 0x6B, 0x81, 0x69, 0x13, 0x5F},
			},
			[]byte{0x5D, 0x22, 0xDF, 0x76},
			false,
		},
		{"dataupcnf",
			args{
				payload: []byte{
					0x80, 0xbd, 0x1f, 0x52, 0x48, 0x80, 0x04, 0x00, 0x05, 0x49, 0x00, 0x00, 0x00, 0x00,
				},
				key: []byte{0x7E, 0x49, 0x1B, 0x08, 0xF3, 0x09, 0xF0, 0x0B, 0xA7, 0xF2, 0xEE, 0x6B, 0x81, 0x69, 0x13, 0x5F},
			},
			[]byte{0xE4, 0xDC, 0x16, 0x81},
			false,
		},
		{"datadownuncnf",
			args{
				payload: []byte{
					0x60, 0xbd, 0x1f, 0x52, 0x48, 0xa0, 0x02, 0x00, 0xd9, 0x14, 0x38, 0x70,
				},
				key: []byte{0x7E, 0x49, 0x1B, 0x08, 0xF3, 0x09, 0xF0, 0x0B, 0xA7, 0xF2, 0xEE, 0x6B, 0x81, 0x69, 0x13, 0x5F},
			},
			[]byte{0xd9, 0x14, 0x38, 0x70},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.args.payload, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calculate() = %v, want %v", hex.EncodeToString(got), hex.EncodeToString(tt.want))
			}
		})
	}
}
