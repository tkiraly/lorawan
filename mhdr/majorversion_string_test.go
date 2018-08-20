package mhdr

import "testing"

func TestMajorVersion_String(t *testing.T) {
	tests := []struct {
		name string
		i    MajorVersion
		want string
	}{
		{"basic",
			MajorVersion(0x00),
			"LoRaWANR1MajorVersion",
		},
		{"basic",
			MajorVersion(0xFF),
			"MajorVersion(255)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("MajorVersion.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
