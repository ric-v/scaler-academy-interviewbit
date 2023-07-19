package arematricessame

import "testing"

func TestAreMatricesSame(t *testing.T) {
	type args struct {
		A [][]int
		B [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, 1},
		{"test2", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 10}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreMatricesSame(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("AreMatricesSame() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkAreMatricesSame-8   	57353151	        24.00 ns/op	       0 B/op	       0 allocs/op
func BenchmarkAreMatricesSame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AreMatricesSame([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	}
}
