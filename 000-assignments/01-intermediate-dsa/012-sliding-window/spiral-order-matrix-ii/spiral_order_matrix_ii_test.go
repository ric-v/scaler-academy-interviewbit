package spiralordermatrixii

import (
	"reflect"
	"testing"
)

func TestSpiralOrderMatrixII(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{A: 3}, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{"test2", args{A: 5}, [][]int{{1, 2, 3, 4, 5}, {16, 17, 18, 19, 6}, {15, 24, 25, 20, 7}, {14, 23, 22, 21, 8}, {13, 12, 11, 10, 9}}},
		{"test3", args{A: 1}, [][]int{{1}}},
		{"test4", args{A: 2}, [][]int{{1, 2}, {4, 3}}},
		{"test5", args{A: 8}, [][]int{{1, 2, 3, 4, 5, 6, 7, 8}, {28, 29, 30, 31, 32, 33, 34, 9}, {27, 48, 49, 50, 51, 52, 35, 10}, {26, 47, 60, 61, 62, 53, 36, 11}, {25, 46, 59, 64, 63, 54, 37, 12}, {24, 45, 58, 57, 56, 55, 38, 13}, {23, 44, 43, 42, 41, 40, 39, 14}, {22, 21, 20, 19, 18, 17, 16, 15}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpiralOrderMatrixII(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpiralOrderMatrixII() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSpiralOrderMatrixII-8   	 1881385	       620.7 ns/op	    1040 B/op	      11 allocs/op
func BenchmarkSpiralOrderMatrixII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SpiralOrderMatrixII(10)
	}
}
