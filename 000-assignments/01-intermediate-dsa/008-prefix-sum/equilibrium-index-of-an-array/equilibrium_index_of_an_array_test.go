package equilibriumindexofanarray

import "testing"

func TestEquilibriumIndex(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{-7, 1, 5, 2, -4, 3, 0}}, 3},
		{"test2", args{[]int{1, 2, 3}}, -1},
		{"test3", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EquilibriumIndex(tt.args.A); got != tt.want {
				t.Errorf("EquilibriumIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkEquilibriumIndex-8   	125865399	         9.178 ns/op	       0 B/op	       0 allocs/op
func BenchmarkEquilibriumIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EquilibriumIndex([]int{-7, 1, 5, 2, -4, 3, 0})
	}
}
