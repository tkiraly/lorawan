package mhdr

import "testing"

func TestMType_String(t *testing.T) {
	tests := []struct {
		name string
		i    MType
		want string
	}{
		{"basic",
			MType(ConfirmedDataDownMessageType),
			"ConfirmedDataDownMessageType",
		},
		{"basic",
			MType(0xFF),
			"MType(255)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("MType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
