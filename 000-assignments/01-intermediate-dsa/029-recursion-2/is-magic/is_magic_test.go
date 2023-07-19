package ismagic

import "testing"

func TestIsMagic(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{83557}, 1},
		{"test2", args{1291}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMagic(tt.args.A); got != tt.want {
				t.Errorf("IsMagic() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkIsMagic-8   	45291752	        25.69 ns/op	       0 B/op	       0 allocs/op
func BenchmarkIsMagic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsMagic(83557)
	}
}
