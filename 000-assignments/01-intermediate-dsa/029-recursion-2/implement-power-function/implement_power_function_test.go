package implementpowerfunction

import "testing"

func TestPowerFunction(t *testing.T) {
	type args struct {
		A int
		B int
		C int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{2, 3, 3}, 2},
		{"test2", args{3, 3, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PowerFunction(tt.args.A, tt.args.B, tt.args.C); got != tt.want {
				t.Errorf("ImplementPowerFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkPowerFunction-8   	27721084	        43.01 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPowerFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PowerFunction(2, 3, 3)
	}
}
