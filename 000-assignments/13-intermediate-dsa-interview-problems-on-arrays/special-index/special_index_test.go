package specialindex

import "testing"

func TestSpecialIndex(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 1, 6, 4}}, 1},
		{"test2", args{[]int{1, 1, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpecialIndex(tt.args.A); got != tt.want {
				t.Errorf("SpecialIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSpecialIndex-8   	18880926	        73.80 ns/op	      64 B/op	       2 allocs/op
func BenchmarkSpecialIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SpecialIndex([]int{2, 1, 6, 4})
	}
}
