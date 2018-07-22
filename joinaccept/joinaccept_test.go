package joinaccept

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		args    args
		want    joinAccept
		wantErr bool
	}{
		{"first",
			args{
				[]byte{0x20, 0xa3, 0x28, 0x30, 0xdf, 0x17, 0xe2, 0x8a,
					0x2a, 0x4c, 0xc9, 0x35, 0x6b, 0x58, 0x71, 0xb8, 0x94,
					0x00, 0xe0, 0x78, 0x4e, 0x1e, 0xcc, 0x10, 0x3f, 0x03,
					0x4d, 0xac, 0x6c, 0x8e, 0x1d, 0x7c, 0xb6}},
			joinAccept([]byte{0x20, 0xa3, 0x28, 0x30, 0xdf, 0x17, 0xe2, 0x8a,
				0x2a, 0x4c, 0xc9, 0x35, 0x6b, 0x58, 0x71, 0xb8, 0x94,
				0x00, 0xe0, 0x78, 0x4e, 0x1e, 0xcc, 0x10, 0x3f, 0x03,
				0x4d, 0xac, 0x6c, 0x8e, 0x1d, 0x7c, 0xb6}),
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
