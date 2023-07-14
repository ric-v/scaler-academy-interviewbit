package shaggyanddistances

import "testing"

func TestShaggyAndDistances(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{7, 1, 3, 4, 1, 7}}, 3},
		{"test2", args{[]int{1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShaggyAndDistances(tt.args.A); got != tt.want {
				t.Errorf("ShaggyAndDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkShaggyAndDistances-8   	11731798	       103.8 ns/op	       0 B/op	       0 allocs/op
func BenchmarkShaggyAndDistances(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShaggyAndDistances([]int{7, 1, 3, 4, 1, 7})
	}
}
