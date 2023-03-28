package besttimetobuyandsellstocks1

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BestTimeToBuyAndSellStocks(tt.args.A); got != tt.want {
				t.Errorf("BestTimeToBuyAndSellStocks() = %v, want %v", got, tt.want)
			}
		})
	}
}
