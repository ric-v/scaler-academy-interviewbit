package sumofdigits

import "testing"

func TestSumOfDigits(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{123}, 6},
		{"test2", args{1234}, 10},
		{"test3", args{12345}, 15},
		{"test4", args{9999}, 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfDigits(tt.args.A); got != tt.want {
				t.Errorf("SumOfDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSumOfDigits-8   	100000000	        13.29 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSumOfDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfDigits(12345)
	}
}
