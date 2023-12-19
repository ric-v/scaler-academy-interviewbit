package verylargepower

import "testing"

func TestVeryLargePower(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1, 1}, 1},
		{"2", args{2, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VeryLargePower(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("VeryLargePower() = %v, want %v", got, tt.want)
			}
		})
	}
}
