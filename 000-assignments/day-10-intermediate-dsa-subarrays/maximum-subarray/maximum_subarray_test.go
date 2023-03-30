package maximumsubarray

import "testing"

func TestMaximumSubArray(t *testing.T) {
	type args struct {
		A int
		B int
		C []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{5, 12, []int{2, 1, 3, 4, 5}}, 12},
		{"test2", args{3, 1, []int{2, 2, 2}}, 0},
		{"test3", args{9, 78, []int{7, 1, 8, 5, 8, 5, 3, 3, 5}}, 45},
		{"test4", args{9, 14, []int{1, 2, 4, 4, 5, 5, 5, 7, 5}}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumSubArray(tt.args.A, tt.args.B, tt.args.C); got != tt.want {
				t.Errorf("MaximumSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkMaximumSubArray-8   	 9836878	       117.8 ns/op	      80 B/op	       1 allocs/op
func BenchmarkMaximumSubArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaximumSubArray(9, 14, []int{1, 2, 4, 4, 5, 5, 5, 7, 5})
	}
}
