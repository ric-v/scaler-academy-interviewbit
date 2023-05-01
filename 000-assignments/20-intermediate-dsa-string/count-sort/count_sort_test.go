package countsort

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
		{"test2", args{[]int{1, 3, 1}}, []int{1, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSort(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkCountSort-8   	13620867	        88.23 ns/op	     128 B/op	       2 allocs/op
func BenchmarkCountSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountSort([]int{1, 4, 1, 2, 7, 5, 2})
	}
}
