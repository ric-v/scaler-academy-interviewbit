package besttimetobuyandsellstocksi

import "testing"

func TestBestTimeToBuyAndSellStocks(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2}}, 1},
		{"test2", args{[]int{1, 4, 5, 2, 4}}, 4},
		{"test3", args{[]int{5, 1, 4, 2}}, 3},
		{"test4", args{[]int{3, 5, 2, 6, 5, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BestTimeToBuyAndSellStocks(tt.args.A); got != tt.want {
				t.Errorf("BestTimeToBuyAndSellStocks() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkBestTimeToBuyAndSellStocks-8   	195003870	         6.532 ns/op	       0 B/op	       0 allocs/op
func BenchmarkBestTimeToBuyAndSellStocks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BestTimeToBuyAndSellStocks([]int{3, 5, 2, 6, 5, 1})
	}
}
