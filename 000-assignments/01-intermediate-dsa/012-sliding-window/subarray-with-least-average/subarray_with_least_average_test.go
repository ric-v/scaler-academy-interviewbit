package subarraywithleastaverage

import "testing"

func TestSubarrayWithLeastAverage(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{3, 7, 90, 20, 10, 50, 40}, 3}, 3},
		{"test2", args{[]int{3, 7, 5, 20, -10, 0, 12}, 2}, 4},
		{"test3", args{[]int{18, 11, 16, 19, 11, 9, 8, 15, 3, 10, 9, 20, 1, 19}, 1}, 12},
		{"test4", args{[]int{20, 3, 13, 5, 10, 14, 8, 5, 11, 9, 1, 11}, 9}, 3},
		{"test5", args{[]int{15, 7, 11, 7, 9, 8, 18, 1, 16, 18, 6, 1, 1, 4, 18}, 6}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubarrayWithLeastAverage(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("SubarrayWithLeastAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSubarrayWithLeastAverage-8   	169960833	         8.189 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSubarrayWithLeastAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubarrayWithLeastAverage([]int{3, 7, 90, 20, 10, 50, 40}, 3)
	}
}
