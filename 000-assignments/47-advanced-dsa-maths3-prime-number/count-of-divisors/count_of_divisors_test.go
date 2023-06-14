package countofdivisors

import (
	"reflect"
	"testing"
)

func TestCountOfDivisors(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{2, 3, 4, 5}}, []int{2, 2, 3, 2}},
		{"test2", args{[]int{8, 9, 10}}, []int{4, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountOfDivisors(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountOfDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCountOfDivisors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountOfDivisors([]int{2, 3, 4, 5})
	}
}