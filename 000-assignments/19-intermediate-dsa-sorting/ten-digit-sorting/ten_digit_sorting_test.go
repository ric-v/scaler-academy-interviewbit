package tendigitsorting

import (
	"reflect"
	"testing"
)

func TestTenDigitSorting(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{15, 11, 7, 19}}, []int{7, 19, 15, 11}},
		{"test2", args{[]int{2, 24, 22, 19}}, []int{2, 19, 24, 22}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TenDigitSorting(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TenDigitSorting() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkTenDigitSorting-8   	 7282586	       169.2 ns/op	      88 B/op	       3 allocs/op
func BenchmarkTenDigitSorting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TenDigitSorting([]int{15, 11, 7, 19})
	}
}
