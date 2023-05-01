package commonelements

import (
	"reflect"
	"sort"
	"testing"
)

func TestCommonElements(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 2, 1}, []int{2, 3, 1, 2}}, []int{1, 2, 2}},
		{"test2", args{[]int{2, 1, 4, 10}, []int{3, 6, 2, 10, 10}}, []int{2, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CommonElements(tt.args.A, tt.args.B)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommonElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkCommonElements-8   	 6231286	       189.6 ns/op	      56 B/op	       3 allocs/op
func BenchmarkCommonElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CommonElements([]int{1, 2, 2, 1}, []int{2, 3, 1, 2})
	}
}
