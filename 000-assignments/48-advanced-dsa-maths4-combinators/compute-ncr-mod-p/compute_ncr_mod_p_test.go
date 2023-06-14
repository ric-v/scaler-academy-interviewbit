package computencrmodp

import "testing"

func TestComputenCrModP(t *testing.T) {
	type args struct {
		A int
		B int
		C int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{5, 2, 13}, 10},
		{"test2", args{6, 2, 13}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComputenCrModP(tt.args.A, tt.args.B, tt.args.C); got != tt.want {
				t.Errorf("ComputenCrModP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkComputenCrModP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ComputenCrModP(5, 2, 13)
	}
}
