package findallprimes

import (
	"reflect"
	"testing"
)

func TestFindAllPrimes(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{7}, []int{2, 3, 5, 7}},
		{"test2", args{12}, []int{2, 3, 5, 7, 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAllPrimes(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAllPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}

//    20510             59320 ns/op          107128 B/op         13 allocs/op
func BenchmarkFindAllPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindAllPrimes(10000)
	}
}
