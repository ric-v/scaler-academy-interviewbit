package singlenumberadvanced

import "testing"

func TestSingleNumberAdvanced(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 2, 3, 1}}, 3},
		{"test2", args{[]int{1, 2, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleNumberAdvanced(tt.args.A); got != tt.want {
				t.Errorf("SingleNumberAdvanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSingleNumberAdvanced-8   	445062913	         2.804 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSingleNumberAdvanced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SingleNumberAdvanced([]int{1, 2, 2, 3, 1})
	}
}
