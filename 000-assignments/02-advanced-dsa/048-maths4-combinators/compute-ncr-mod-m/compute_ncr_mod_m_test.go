package computencrmodm

import "testing"

func TestComputeNcrModM(t *testing.T) {
	type args struct {
		n int
		r int
		m int
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
			if got := ComputeNcrModM(tt.args.n, tt.args.r, tt.args.m); got != tt.want {
				t.Errorf("ComputeNcrModM() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkComputeNcrModM-8   	       1	42071562717 ns/op	       0 B/op	       0 allocs/op
func BenchmarkComputeNcrModM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ComputeNcrModM(1000000000, 1000000000, 1000000000)
	}
}
