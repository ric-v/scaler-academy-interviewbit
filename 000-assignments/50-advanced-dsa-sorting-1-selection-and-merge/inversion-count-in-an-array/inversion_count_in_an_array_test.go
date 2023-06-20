package inversioncountinanarray

import "testing"

func TestInversionCountInAnArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 4, 1, 3, 5}}, 3},
		{"test2", args{[]int{2, 3, 4, 5, 6}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InversionCountInAnArray(tt.args.A); got != tt.want {
				t.Errorf("InversionCountInAnArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkInversionCountInAnArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InversionCountInAnArray([]int{2, 4, 1, 3, 5})
	}
}
