package goodpair

import "testing"

func TestGoodPair(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test 1", args: args{A: []int{1, 2, 3, 4}, B: 7}, want: 1},
		{name: "Test 2", args: args{A: []int{1, 2, 4}, B: 4}, want: 0},
		{name: "Test 3", args: args{A: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, B: 10}, want: 1},
		{name: "Test 4", args: args{A: []int{1, 2, 2}, B: 4}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GoodPair(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("GoodPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkGoodPair-8   	61216369	        20.13 ns/op	       0 B/op	       0 allocs/op
func BenchmarkGoodPair(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// input of len 100
		GoodPair([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10)
	}
}
