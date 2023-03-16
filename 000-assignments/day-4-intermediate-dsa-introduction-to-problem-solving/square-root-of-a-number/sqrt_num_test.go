package numsqrt

import "testing"

func TestSqrt(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{A: 4}, 2},
		{"Test 2", args{A: 11}, -1},
		{"Test 3", args{A: 9}, 3},
		{"Test 4", args{A: 0}, 0},
		{"Test 5", args{A: 1}, 1},
		{"Test 6", args{A: 2}, -1},
		{"Test 7", args{A: 3}, -1},
		{"Test 8", args{A: 5}, -1},
		{"Test 9", args{A: 6}, -1},
		{"Test 10", args{A: 100}, 10},
		{"Test 11", args{A: 101}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sqrt(tt.args.A); got != tt.want {
				t.Errorf("Sqrt(%v) = %v, want %v", tt.args.A, got, tt.want)
			}
		})
	}
}
