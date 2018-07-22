package joinaccept

import (
	"reflect"
	"testing"

	"github.com/tkiraly/lorawan/mhdr"
)

func TestParse(t *testing.T) {
	type args struct {
		p   []byte
		key []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *joinAccept
		wantErr bool
	}{
		{"first",
			args{
				[]byte{0x20, 0xa3, 0x28, 0x30, 0xdf, 0x17, 0xe2, 0x8a,
					0x2a, 0x4c, 0xc9, 0x35, 0x6b, 0x58, 0x71, 0xb8, 0x94,
					0x00, 0xe0, 0x78, 0x4e, 0x1e, 0xcc, 0x10, 0x3f, 0x03,
					0x4d, 0xac, 0x6c, 0x8e, 0x1d, 0x7c, 0xb6},
				[]byte{0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02,
					0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02, 0x02}},
			&joinAccept{
				MHDR:     mhdr.MHDR{MType: mhdr.JoinAcceptMessage, Major: mhdr.LoRaWANR1Version},
				AppNonce: []byte{0x24, 0x63, 0x61},
				NetID:    []byte{0x00, 0x00, 0x24},
				DevAddr:  []byte{0x48, 0x3A, 0x24, 0x9D},
				DlSettings: joinaccept.DlSettings{
					RX1DRoffset: 0,
					RX2Datarate: 0,
				},
				RxDelay: 0x00,
				MIC:     []byte{0x5A, 0x8D, 0x60, 0x3A},
				CfList: &joinaccept.CFList{
					Ch4: 8671000,
					Ch5: 8673000,
					Ch6: 8675000,
					Ch7: 8677000,
					Ch8: 8679000,
				},
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := joinaccept.Parse(tt.args.p, tt.args.key)
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
