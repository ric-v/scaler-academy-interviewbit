package pubg

import "testing"

func TestPUBG(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{6, 4}}, 2},
		{"test2", args{[]int{2, 3, 4}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PUBG(tt.args.A); got != tt.want {
				t.Errorf("PUBG() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkPUBG-8   	54038232	        22.89 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPUBG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PUBG([]int{6, 4})
	}
}
