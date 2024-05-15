package iseven

import "testing"

func TestIsEven(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{1}, false},
		{"2", args{2}, true},
		{"3", args{3}, false},
		{"4", args{4}, true},
		{"5", args{5}, false},
		{"6", args{6}, true},
		{"7", args{7}, false},
		{"8", args{8}, true},
		{"9", args{9}, false},
		{"0", args{0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args.n); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
