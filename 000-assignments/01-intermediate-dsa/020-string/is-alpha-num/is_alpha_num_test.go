package isalphanum

import "testing"

func TestIsAlNum(t *testing.T) {
	type args struct {
		A []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]byte{'a', 'b', 'c'}}, 1},
		{"test2", args{[]byte{'a', 'b', 'c', '0'}}, 1},
		{"test3", args{[]byte{'a', 'b', 'c', '0', '1', '2'}}, 1},
		{"test4", args{[]byte{'a', 'b', 'c', '0', '1', '2', 'A', 'B', 'C'}}, 1},
		{"test5", args{[]byte{'a', 'b', 'c', '0', '1', '2', 'A', 'B', 'C', ' ', '-', '.'}}, 0},
		{"test6", args{[]byte{'a', 'b', 'c', '0', '1', '2', 'A', 'B', 'C', ' ', '-', '.', 'a', 'b', 'c', '0', '1', '2', 'A', 'B', 'C', ' ', '-', '.'}}, 0},
		{"test7", args{[]byte{'a', 'b', 'c', '0', '1', '2', 'A', 'B', 'C', ' ', '-', '.', 'a', 'b', 'c', '0', '1', '2', 'A', 'B', 'C', ' ', '-', '.', 'a', 'b', 'c', '0', '1', '2', 'A', 'B', 'C', ' ', '-', '.'}}, 0},
		{"test8", args{[]byte{'a', 'b', 'c', '0', '1', '2', 'A', 'B', 'C', ' ', '-', '.', 'a', 'b', 'c', '0', '1', '2', 'A', 'B', 'C', ' ', '-', '.', 'a', 'b', 'c', '0', '1', '2', 'A', 'B', 'C', ' ', '-', '.', 'a', 'b', 'c', '0', '1', '2', 'A', 'B', 'C', ' ', '-', '.'}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAlNum(tt.args.A); got != tt.want {
				t.Errorf("IsAlNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkIsAlNum-8   	132629154	         9.606 ns/op	       0 B/op	       0 allocs/op
func BenchmarkIsAlNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsAlNum([]byte{'a', 'b', 'c', '0', '1', '2', 'A', 'B', 'C', ' ', '-', '.'})
	}
}
