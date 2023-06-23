package countsort2

import (
	"reflect"
	"testing"
)

func TestCountSort(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 4, 1, 2, 7, 5, 2}}, []int{1, 1, 2, 2, 4, 5, 7}},
		{"test2", args{[]int{1, 4, 1, 2, 7, 5, 2, 0}}, []int{0, 1, 1, 2, 2, 4, 5, 7}},
		{"test3", args{[]int{1, 4, 1, 2, 7, 5, 2, 0, 0}}, []int{0, 0, 1, 1, 2, 2, 4, 5, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSort(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCountSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountSort([]int{1, 4, 1, 2, 7, 5, 2, 0, 0})
	}
}
