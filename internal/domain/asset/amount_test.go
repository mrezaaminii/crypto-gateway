package asset

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		input string
		decimals int
		want string
		wantErr bool
	} {
		{"1.5", 18, "1500000000000000000", false},
		{"0.000001", 6, "1", false},
		{"100", 8, "10000000000", false},
		{"1.1234567", 6, "", true},
		{"-1", 18, "", true},
		{"abc", 18, "", true},
	}

	for _, tt := range tests {
		got, err := ParseAmount(tt.input, tt.decimals)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseAmount(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if !tt.wantErr && got.value.String() != tt.want {
			t.Errorf("ParseAmount(%q) = %s, want %s", tt.input, got.value, tt.want)
		}
	}
}