package dungeonprincess

import "testing"

func TestDungeonPrincess(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}}, 7},
		{"test2", args{[][]int{{1, -1, 0}, {-1, 1, -1}, {1, 0, -1}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DungeonPrincess(tt.args.A); got != tt.want {
				t.Errorf("DungeonPrincess() = %v, want %v", got, tt.want)
			}
		})
	}
}
