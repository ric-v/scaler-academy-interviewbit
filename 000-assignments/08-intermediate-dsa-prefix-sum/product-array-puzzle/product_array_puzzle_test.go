package productarraypuzzle

import (
	"reflect"
	"testing"
)

func TestProductArrayPuzzle(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5}}, []int{120, 60, 40, 30, 24}},
		{"test2", args{[]int{5, 1, 10, 1}}, []int{10, 50, 5, 50}},
		{"test_large", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}}, []int{2432902008176640000, 1216451004088320000, 810967336058880000, 608225502044160000, 486580401635328000, 405483668029440000, 347557429739520000, 304112751022080000, 270322445352960000, 243290200817664000, 221172909834240000, 202741834014720000, 187146308321280000, 173778714869760000, 162193467211776000, 152056375511040000, 143111882833920000, 135161222676480000, 128047474114560000, 121645100408832000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProductArrayPuzzle(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductArrayPuzzle() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkProductArrayPuzzle-8   	 4961422	       223.1 ns/op	       0 B/op	       0 allocs/op
func BenchmarkProductArrayPuzzle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductArrayPuzzle([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	}
}
