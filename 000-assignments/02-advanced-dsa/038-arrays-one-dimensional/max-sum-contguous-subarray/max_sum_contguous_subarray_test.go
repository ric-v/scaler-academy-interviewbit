package maxsumcontguoussubarray

import "testing"

func TestMaxSumContiguousSubarray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 4, -10}}, 10},
		{"test2", args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}}, 6},
		{"test3", args{[]int{-2, -3, 4, -1, -2, 1, 5, -3}}, 7},
		{"test4", args{[]int{-2, -3, -4, -1, -2, -1, -5, -3}}, -1},
		{"testlarge", args{[]int{-2, -3, -4, -1, -2, -1, -5, -3, -2, -3, -4, -1, -2, -1, -5, -3, -2, -3, -4, -1, -2, -1, -5, -3, -2, -3, -4, -1, -2, -1, -5, -3, -2, -3, -4, -1, -2, -1, -5, -3, -2, -3, -4, -1, -2, -1, -5, -3, -2, -3, -4, -1, -2, -1, -5, -3, -2, -3, -4, -1, -2, -1, -5, -3}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSumContiguousSubarray(tt.args.A); got != tt.want {
				t.Errorf("MaxSumContiguousSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMaxSumContiguousSubarray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxSumContiguousSubarray([]int{-2, -3, -4, -1, -2, -1, -5, -3})
	}
}
