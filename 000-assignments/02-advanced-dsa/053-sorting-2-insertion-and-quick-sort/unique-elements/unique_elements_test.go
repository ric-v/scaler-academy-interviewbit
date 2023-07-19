package uniqueelements

import "testing"

func TestUniqueElements(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 1, 3}}, 1},
		{"test2", args{[]int{2, 4, 5}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueElements(tt.args.A); got != tt.want {
				t.Errorf("UniqueElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkUniqueElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniqueElements([]int{1, 1, 3})
	}
}
