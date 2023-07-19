package maximumsubmatrixsum

import "testing"

func TestMaximumSubmatrixSum(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test1", args{[][]int{{-5, -4, -3}, {-1, 2, 3}, {2, 2, 4}}}, 12},
		{"test2", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, 45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumSubmatrixSum(tt.args.A); got != tt.want {
				t.Errorf("MaximumSubmatrixSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMaximumSubmatrixSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaximumSubmatrixSum([][]int{{-5, -4, -3}, {-1, 2, 3}, {2, 2, 4}})
	}
}
