package rotatedsortedarraysearch

import "testing"

func TestRotatedSortedArraySearch(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 4}, 0},
		{"test2", args{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 5}, 1},
		{"test3", args{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 6}, 2},
		{"test4", args{[]int{4, 5, 6, 7, 0, 1, 2, 3}, 2}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotatedSortedArraySearch(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("RotatedSortedArraySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkRotatedSortedArraySearch-8   	174352030	         6.736 ns/op	       0 B/op	       0 allocs/op
func BenchmarkRotatedSortedArraySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RotatedSortedArraySearch([]int{4, 5, 6, 7, 0, 1, 2, 3}, 4)
	}
}
