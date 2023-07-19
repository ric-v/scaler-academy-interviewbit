package subarraywith0sum

import "testing"

func TestSubArrayWith0Sum(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5, 6}}, 0},
		{"test2", args{[]int{1, 2, 3, 4, -4, 6}}, 1},
		{"test3", args{[]int{1, -1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubArrayWith0Sum(tt.args.A); got != tt.want {
				t.Errorf("SubArrayWith0Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSubArrayWith0Sum-8   	 5141236	       250.1 ns/op	     112 B/op	       3 allocs/op
func BenchmarkSubArrayWith0Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubArrayWith0Sum([]int{1, 2, 3, 4, 5, 6})
	}
}
