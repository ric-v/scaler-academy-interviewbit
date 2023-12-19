package primemoduloinverse

import "testing"

func TestPrimeModuloInverse(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3, 5}, 2},
		{"2", args{6, 23}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrimeModuloInverse(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("PrimeModuloInverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
