package pairswithgivensumii

import "testing"

func TestPairsWithGivenSumII(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 1, 1}, 2}, 3},
		{"test2", args{[]int{1, 5, 7, 10}, 8}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PairsWithGivenSumII(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("PairsWithGivenSumII() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkPairsWithGivenSumII-8   	281759580	         4.443 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPairsWithGivenSumII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PairsWithGivenSumII([]int{1, 1, 1}, 2)
	}
}
