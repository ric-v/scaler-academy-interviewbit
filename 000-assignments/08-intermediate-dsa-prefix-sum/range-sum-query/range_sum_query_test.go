package rangesumquery

import (
	"reflect"
	"testing"
)

func TestRangeSumQuery(t *testing.T) {
	type args struct {
		A []int
		B [][]int
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5}, [][]int{{0, 3}, {1, 2}}}, []int64{10, 5}},
		{"test2", args{[]int{2, 2, 2}, [][]int{{0, 0}, {1, 2}}}, []int64{2, 4}},
		{"test_large", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, [][]int{{0, 3}, {1, 2}, {0, 9}, {0, 0}, {9, 9}}}, []int64{10, 5, 55, 1, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RangeSumQuery(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RangeSumQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkRangeSumQuery-8   	 8380905	       143.7 ns/op	     120 B/op	       4 allocs/op
func BenchmarkRangeSumQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeSumQuery([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, [][]int{{0, 3}, {1, 2}, {0, 9}, {0, 0}, {9, 9}})
	}
}
