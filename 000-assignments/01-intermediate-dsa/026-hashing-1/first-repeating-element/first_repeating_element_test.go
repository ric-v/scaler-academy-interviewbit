package firstrepeatingelement

import "testing"

func TestFirstRepeatingElement(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{10, 5, 3, 4, 3, 5, 6}}, 5},
		{"test2", args{[]int{6, 10, 5, 4, 9, 120}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstRepeatingElement(tt.args.A); got != tt.want {
				t.Errorf("FirstRepeatingElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkFirstRepeatingElement-8   	 7507834	       156.0 ns/op	       0 B/op	       0 allocs/op
func BenchmarkFirstRepeatingElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstRepeatingElement([]int{10, 5, 3, 4, 3, 5, 6})
	}
}
