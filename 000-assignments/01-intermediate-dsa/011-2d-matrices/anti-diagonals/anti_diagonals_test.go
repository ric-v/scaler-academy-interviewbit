package antidiagonals

import (
	"reflect"
	"testing"
)

func TestAntiDiagonals(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, [][]int{{1, 0, 0}, {2, 4, 0}, {3, 5, 7}, {6, 8, 0}, {9, 0, 0}}},
		{"test2", args{[][]int{{1, 2}, {3, 4}}}, [][]int{{1, 0}, {2, 3}, {4, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AntiDiagonals(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AntiDiagonals() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkAntiDiagonals-8   	 1427215	       878.4 ns/op	     640 B/op	      19 allocs/op
func BenchmarkAntiDiagonals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AntiDiagonals([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	}
}
