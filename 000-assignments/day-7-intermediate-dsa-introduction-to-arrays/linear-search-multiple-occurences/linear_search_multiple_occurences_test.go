package linearsearchpmultipleoccurences

import "testing"

func TestLinearSearchMultipleOccurences(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5}, 1}, 1},
		{"test2", args{[]int{5, 4, 3, 2, 1}, 1}, 1},
		{"test3", args{[]int{1, 1, 1, 1, 1}, 1}, 5},
		{"test4", args{[]int{8, 1, 2, 2, 3, 8, 4, 5, 7, 6}, 8}, 2},
		{"test5", args{[]int{-2, 1, -4, 5, 3}, 1}, 1},
		{"test6", args{[]int{1, 2, 2}, 2}, 2},
		{"test7", args{[]int{1, 2, 1}, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinearSearchMultipleOccurences(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("LinearSearchMultipleOccurences() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkLinearSearchMultipleOccurences-8   	24349117	        49.42 ns/op	       0 B/op	       0 allocs/op
func BenchmarkLinearSearchMultipleOccurences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LinearSearchMultipleOccurences([]int{8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6, 8, 1, 2, 2, 3, 8, 4, 5, 7, 6}, 8)
	}
}