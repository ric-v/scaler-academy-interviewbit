package powerwithmodules

import "testing"

func TestPowerWithModules(t *testing.T) {
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
		{"test2", args{5, 2, 4}, 1},
		{"test3", args{76277, 376, 2417}, 2254},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PowerWithModules(tt.args.A, tt.args.B, tt.args.C); got != tt.want {
				t.Errorf("PowerWithModules() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkPowerWithModules-8   	 9804922	       120.8 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPowerWithModules(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PowerWithModules(76277, 376, 2417)
	}
}
