package countsubarrayzerosum

import "testing"

func TestCountSubarrayZeroSum(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, -1, -2, 2}}, 3},
		{"test2", args{[]int{-1, 2, -1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSubarrayZeroSum(tt.args.A); got != tt.want {
				t.Errorf("CountSubarrayZeroSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkCountSubarrayZeroSum-8   	 6228061	       195.5 ns/op	      48 B/op	       2 allocs/op
func BenchmarkCountSubarrayZeroSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountSubarrayZeroSum([]int{1, -1, -2, 2})
	}
}
