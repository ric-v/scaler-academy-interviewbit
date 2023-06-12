package singlenumberadvancediii

import (
	"reflect"
	"testing"
)

func TestSingleNumberAdvancedIII(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 1, 2, 4}}, []int{3, 4}},
		{"test2", args{[]int{1, 2}}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleNumberAdvancedIII(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SingleNumberAdvancedIII() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSingleNumberAdvancedIII-8   	35970898	        33.51 ns/op	      16 B/op	       1 allocs/op
func BenchmarkSingleNumberAdvancedIII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SingleNumberAdvancedIII([]int{1, 2, 3, 1, 2, 4})
	}
}
