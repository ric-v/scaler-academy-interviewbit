package factorssort

import (
	"reflect"
	"testing"
)

func TestFactorSort(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{6, 8, 9}}, []int{9, 6, 8}},
		{"test2", args{[]int{2, 4, 7}}, []int{2, 7, 4}},
		{"test3", args{[]int{36, 13, 13, 26, 37, 28, 27, 43, 7}}, []int{7, 13, 13, 37, 43, 26, 27, 28, 36}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FactorSort(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FactorSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkFactorSort-8   	 3300835	       361.7 ns/op	      80 B/op	       3 allocs/op
func BenchmarkFactorSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorSort([]int{6, 8, 9})
	}
}
