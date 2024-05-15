package iseven

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{"0", 0, true},
		{"1", 1, false},
		{"2", 2, true},
		{"3", 3, false},
		{"4", 4, true},
		{"5", 5, false},
		{"6", 6, true},
		{"7", 7, false},
		{"8", 8, true},
		{"9", 9, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
