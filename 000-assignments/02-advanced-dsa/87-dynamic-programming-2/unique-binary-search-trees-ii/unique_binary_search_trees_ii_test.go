package uniquebinarysearchtreesii

import "testing"

func TestUniqueBinarySearchTreesII(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{1}, 1},
		{"test2", args{2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueBinarySearchTreesII(tt.args.A); got != tt.want {
				t.Errorf("UniqueBinarySearchTreesII() = %v, want %v", got, tt.want)
			}
		})
	}
}
