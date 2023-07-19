package concatenatethreenumbers

import "testing"

func TestConcatenateThreeNumbers(t *testing.T) {
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
		{"test1", args{1, 2, 3}, 123},
		{"test2", args{10, 20, 30}, 102030},
		{"test3", args{55, 43, 47}, 434755},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcatenateThreeNumbers(tt.args.A, tt.args.B, tt.args.C); got != tt.want {
				t.Errorf("ConcatenateThreeNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkConcatenateThreeNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatenateThreeNumbers(1, 2, 3)
	}
}
