package stairs

import "testing"

func TestStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{1}, 1},
		{"test2", args{2}, 2},
		{"test3", args{3}, 3},
		{"test4", args{4}, 5},
		{"test5", args{5}, 8},
		{"test6", args{6}, 13},
		{"test7", args{7}, 21},
		{"test8", args{8}, 34},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Stairs(tt.args.n); got != tt.want {
				t.Errorf("Stairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkStairs-8   	28947606	        43.23 ns/op	      96 B/op	       1 allocs/op
func BenchmarkStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Stairs(10)
	}
}
