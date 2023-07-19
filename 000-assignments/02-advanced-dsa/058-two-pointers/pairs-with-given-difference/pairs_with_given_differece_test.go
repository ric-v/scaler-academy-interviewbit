package pairswithgivendifference

import "testing"

func TestPairsWithGivenDifference(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 5, 3, 4, 2}, 3}, 2},
		{"test2", args{[]int{8, 12, 16, 4, 0, 20}, 4}, 5},
		{"test3", args{[]int{1, 1, 1, 2, 2}, 0}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PairsWithGivenDifference(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("PairsWithGivenDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkPairsWithGivenDifference-8   	10407451	       118.8 ns/op	      72 B/op	       2 allocs/op
func BenchmarkPairsWithGivenDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PairsWithGivenDifference([]int{1, 5, 3, 4, 2}, 3)
	}
}
