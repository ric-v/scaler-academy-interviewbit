package isprime

import "testing"

func TestIsPrime(t *testing.T) {
	type args struct {
		A int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{1},
			0,
		},
		{
			"2",
			args{2},
			1,
		},
		{
			"3",
			args{3},
			1,
		},
		{
			"4",
			args{4},
			0,
		},
		{
			"5",
			args{5},
			1,
		},
		{
			"10",
			args{10},
			0,
		},
		{
			"100",
			args{100},
			0,
		},
		{
			"1000",
			args{1000},
			0,
		},
		{
			"73",
			args{73},
			1,
		},
		{
			"1000000007",
			args{1000000007},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.A); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkIsPrime-8   	    5232	    244470 ns/op	       0 B/op	       0 allocs/op
func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(1000000007)
	}
}
