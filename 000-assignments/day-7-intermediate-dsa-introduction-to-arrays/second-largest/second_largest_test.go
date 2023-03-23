package secondlargest

import "testing"

func TestSecondLargest(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[]int{1, 2, 3, 4, 5}}, 4},
		{"Test 2", args{[]int{5, 4, 3, 2, 1}}, 4},
		{"Test 3", args{[]int{1, 1, 1, 1, 1}}, -1},
		{"Test 4", args{[]int{1, 2, 3, 4, 4}}, 3},
		{"Test 5", args{[]int{2, 1, 2}}, 1},
		{"Test 6", args{[]int{2}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SecondLargest(tt.args.A); got != tt.want {
				t.Errorf("SecondLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
