package anybasetodecimal

import "testing"

func TestAnyBaseToDecimal(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{"test1", args{A: 101, B: 2}, 5},
		{"test2", args{A: 101, B: 3}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := AnyBaseToDecimal(tt.args.A, tt.args.B); gotResult != tt.wantResult {
				t.Errorf("AnyBaseToDecimal() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

// BenchmarkAnyBaseToDecimal-8   	366305275	         3.542 ns/op	       0 B/op	       0 allocs/o
func BenchmarkAnyBaseToDecimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AnyBaseToDecimal(101, 2)
	}
}
