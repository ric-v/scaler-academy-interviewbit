package minnumofsquares

import "testing"

func TestMinNumOfSquares(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{6}, 3},
		{"test2", args{5}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinNumOfSquares(tt.args.A); got != tt.want {
				t.Errorf("MinNumOfSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
