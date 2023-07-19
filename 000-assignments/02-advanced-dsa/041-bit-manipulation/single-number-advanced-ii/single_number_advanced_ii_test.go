package singlenumberadvancedii

import "testing"

func TestSingleNumberAdvancedII(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 4, 3, 3, 2, 2, 3, 1, 1}}, 4},
		{"test2", args{[]int{0, 0, 0, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleNumberAdvancedII(tt.args.A); got != tt.want {
				t.Errorf("SingleNumberAdvancedII() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSingleNumberAdvancedII-8   	 7287859	       155.8 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSingleNumberAdvancedII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SingleNumberAdvancedII([]int{1, 2, 4, 3, 3, 2, 2, 3, 1, 1})
	}
}
