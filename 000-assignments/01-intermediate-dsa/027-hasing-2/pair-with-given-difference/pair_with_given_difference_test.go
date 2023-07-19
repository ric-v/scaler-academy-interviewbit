package pairwithgivendifference

import "testing"

func TestPairWithGivenDifference(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{5, 10, 3, 2, 50, 80}, 78}, 1},
		{"test2", args{[]int{-10, 20}, 30}, 1},
		{"test3", args{[]int{-100, 20, 80}, 9}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PairWithGivenDifference(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("PairWithGivenDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkPairWithGivenDifference-8   	 7860970	       134.7 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPairWithGivenDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PairWithGivenDifference([]int{5, 10, 3, 2, 50, 80}, 78)
	}
}
