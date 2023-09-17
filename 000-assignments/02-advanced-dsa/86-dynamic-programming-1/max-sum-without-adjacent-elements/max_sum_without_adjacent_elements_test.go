package maxsumwithoutadjacentelements

import "testing"

func TestMaxSumWithoutAdjacentElements(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[][]int{{1}, {2}}}, 2},
		{"test2", args{[][]int{{1, 2, 3, 4}, {2, 3, 4, 5}}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSumWithoutAdjacentElements(tt.args.A); got != tt.want {
				t.Errorf("MaxSumWithoutAdjacentElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
