package searchforarange

import (
	"reflect"
	"testing"
)

func TestSearchForARange(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
		{"test2", args{[]int{5, 17, 100, 111}, 3}, []int{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchForARange(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchForARange() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSearchForARange-8   	46608819	        26.37 ns/op	      16 B/op	       1 allocs/op
func BenchmarkSearchForARange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchForARange([]int{5, 7, 7, 8, 8, 10}, 8)
	}
}
