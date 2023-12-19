package pairsumdivisiblebym

import "testing"

func TestPairSumDivisibleByM(t *testing.T) {
	type args struct {
		A []int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 4, 5}, 2}, 4},
		{"2", args{[]int{5, 17, 100, 11}, 28}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PairSumDivisibleByM(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("PairSumDivisibleByM() = %v, want %v", got, tt.want)
			}
		})
	}
}
