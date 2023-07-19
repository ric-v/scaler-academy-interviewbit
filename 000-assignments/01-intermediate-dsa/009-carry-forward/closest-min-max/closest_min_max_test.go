package closestminmax

import "testing"

func TestClosestMinMax(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 2}}, 2},
		{"test2", args{[]int{2, 6, 1, 6, 9}}, 3},
		{"test3", args{[]int{1, 1, 1, 1, 1}}, 1},
		{"test4", args{[]int{8, 1, 2, 2, 3, 8, 4, 5, 7, 6}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClosestMinMax(tt.args.A); got != tt.want {
				t.Errorf("ClosesMinMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkClosestMinMax-8   	41898853	        32.79 ns/op	       0 B/op	       0 allocs/op
func BenchmarkClosestMinMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClosestMinMax([]int{8, 1, 2, 2, 3, 8, 4, 5, 7, 6})
	}
}
