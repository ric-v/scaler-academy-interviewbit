package searchinarowwiseandcolumnwisesortedmatrix

import "testing"

func TestSearchInMatrix(t *testing.T) {
	type args struct {
		A [][]int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"test1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 2}, 1011},
		{"test2", args{[][]int{{1, 2}, {3, 3}}, 3}, 2019},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInMatrix(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("SearchInMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSearchInMatrix-8   	74407530	        14.01 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSearchInMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchInMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 2)
	}
}
