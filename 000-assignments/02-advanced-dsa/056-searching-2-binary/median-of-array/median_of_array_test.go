package medianofarray

import "testing"

func TestMedianOfArray(t *testing.T) {
	type args struct {
		num1 []int
		num2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{[]int{1, 4, 5}, []int{2, 3}}, 3.0},
		{"test2", args{[]int{1, 2, 3}, []int{4}}, 2.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MedianOfArray(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("MedianOfArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMedianOfArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MedianOfArray([]int{1, 4, 5}, []int{2, 2})
	}
}
