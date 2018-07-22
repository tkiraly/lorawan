package joinrequest

import (
	"reflect"
	"testing"

	"github.com/tkiraly/lorawan/mhdr"
)

func TestNew(t *testing.T) {
	type args struct {
		Major    mhdr.MajorVersion
		appeui   []byte
		deveui   []byte
		appkey   []byte
		devnonce []byte
	}
	tests := []struct {
		name    string
		args    args
		want    JoinRequest
		wantErr bool
	}{
		{"basic", args{
			Major:    mhdr.LoRaWANR1MajorVersion,
			appeui:   []byte{0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x22},
			deveui:   []byte{0x33, 0x22, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33},
			appkey:   []byte{0x33, 0x22, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x22, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33},
			devnonce: []byte{0x33, 0x22},
		}, joinRequest([]byte{0x00, 0x22, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x22,
			0x33, 0x22, 0x33, 0x6F, 0x2A, 0x86, 0x4B}),
			false},
		{"geterror", args{
			Major:    mhdr.LoRaWANR1MajorVersion,
			appeui:   []byte{0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x22},
			deveui:   []byte{0x33, 0x22, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33},
			appkey:   []byte{0x33, 0x22, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x22, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33},
			devnonce: []byte{0x33, 0x22},
		}, joinRequest([]byte{0x00, 0x22, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x22,
			0x33, 0x22, 0x33, 0x6F, 0x2A, 0x86, 0x4B}),
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.Major, tt.args.appeui, tt.args.deveui, tt.args.appkey, tt.args.devnonce)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		args    args
		want    JoinRequest
		wantErr bool
	}{
		// TODO: Add test cases.
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
