package maindiagonalsum

import "testing"

func TestMainDiagonalSum(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, 15},
		{"test2", args{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}}, 34},
		{"test3", args{[][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}}, 65},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MainDiagonalSum(tt.args.A); got != tt.want {
				t.Errorf("MainDiagonalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkMainDiagonalSum-8   	57766184	        23.33 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMainDiagonalSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MainDiagonalSum([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}})
	}
}
