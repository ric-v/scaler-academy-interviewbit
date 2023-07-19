package containerwithmostwater

import "testing"

func TestContainerWithMostWater(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 5, 4, 3}}, 6},
		{"test2", args{[]int{1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainerWithMostWater(tt.args.A); got != tt.want {
				t.Errorf("ContainerWithMostWater() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkContainerWithMostWater-8   	271992824	         4.084 ns/op	       0 B/op	       0 allocs/op
func BenchmarkContainerWithMostWater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ContainerWithMostWater([]int{1, 5, 4, 3})
	}
}
