package minordiagonalsum

import "testing"

func TestMinorDiagonalSum(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[][]int{{1, -2, -3}, {-4, 5, -6}, {-7, -8, 9}}}, -5},
		{"test2", args{[][]int{{3, 2}, {2, 3}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinorDiagonalSum(tt.args.A); got != tt.want {
				t.Errorf("MinorDiagonalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkMinorDiagonalSum-8   	148378294	         8.396 ns/op	       0 B/op	       0 allocs/op
func BenchmarkMinorDiagonalSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinorDiagonalSum([][]int{{1, -2, -3}, {-4, 5, -6}, {-7, -8, 9}})
	}
}
