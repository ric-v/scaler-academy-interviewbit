package mergetwosortedarrays

import (
	"reflect"
	"testing"
)

func TestMergeTwoSortedArrays(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 3, 5}, []int{2, 4, 6}}, []int{1, 2, 3, 4, 5, 6}},
		{"test2", args{[]int{1, 3, 5}, []int{2, 4, 6, 7}}, []int{1, 2, 3, 4, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeTwoSortedArrays(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeTwoSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkMergeTwoSortedArrays-8   	 9611456	       125.5 ns/op	     120 B/op	       4 allocs/op
func BenchmarkMergeTwoSortedArrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeTwoSortedArrays([]int{1, 3, 5}, []int{2, 4, 6})
	}
}
