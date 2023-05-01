package frequencyofelementquery

import (
	"reflect"
	"testing"
)

func TestFreqOfElementQuery(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 1, 1}, []int{1, 2}}, []int{3, 1}},
		{"test2", args{[]int{2, 5, 9, 2, 8}, []int{3, 2}}, []int{0, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FreqOfElementQuery(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FreqOfElementQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkFreqOfElementQuery-8   	14204545	        84.33 ns/op	      16 B/op	       1 allocs/op
func BenchmarkFreqOfElementQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FreqOfElementQuery([]int{1, 2, 1, 1}, []int{1, 2})
	}
}
