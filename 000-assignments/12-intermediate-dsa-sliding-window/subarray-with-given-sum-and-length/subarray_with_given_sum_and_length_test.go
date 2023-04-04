package subarraywithgivensumandlength

import "testing"

func TestSubarrayWithGivenSumAndLength(t *testing.T) {
	type args struct {
		A []int
		B int
		C int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{4, 3, 2, 6, 1}, 3, 11}, 1},
		{"test2", args{[]int{4, 2, 2, 5, 1}, 4, 6}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubarrayWithGivenSumAndLength(tt.args.A, tt.args.B, tt.args.C); got != tt.want {
				t.Errorf("SubarrayWithGivenSumAndLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSubarrayWithGivenSumAndLength-8   	267080392	         4.523 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSubarrayWithGivenSumAndLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubarrayWithGivenSumAndLength([]int{4, 3, 2, 6, 1}, 3, 11)
	}
}
