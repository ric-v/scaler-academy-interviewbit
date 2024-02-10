package uniquepathsinagrid

import "testing"

func TestUniquePathsInAGrid(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}}, 2},
		{"test2", args{[][]int{{0, 0, 0}, {1, 1, 1}, {0, 0, 0}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniquePathsInAGrid(tt.args.A); got != tt.want {
				t.Errorf("UniquePathsInAGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
