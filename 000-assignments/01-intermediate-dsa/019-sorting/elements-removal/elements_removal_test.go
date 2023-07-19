package elementsremoval

import "testing"

func TestElementsRemoval(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1 ", args{[]int{2, 1}}, 4},
		{"test2 ", args{[]int{5}}, 5},
		{"test3 ", args{[]int{1, 2, 3, 4, 5}}, 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ElementsRemoval(tt.args.A); got != tt.want {
				t.Errorf("ElementsRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkElementsRemoval-8   	 5465830	       207.9 ns/op	      88 B/op	       3 allocs/op
func BenchmarkElementsRemoval(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ElementsRemoval([]int{1, 2, 3, 4, 5})
	}
}
