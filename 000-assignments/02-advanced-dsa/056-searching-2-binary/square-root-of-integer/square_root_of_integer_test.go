package squarerootofinteger

import "testing"

func TestSquareRootOfInteger(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{11}, 3},
		{"test2", args{9}, 3},
		{"test3", args{50}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SquareRootOfInteger(tt.args.A); got != tt.want {
				t.Errorf("SquareRootOfInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSquareRootOfInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SquareRootOfInteger(11)
	}
}
