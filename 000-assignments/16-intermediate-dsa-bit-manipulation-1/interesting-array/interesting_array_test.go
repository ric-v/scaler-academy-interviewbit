package interestingarray

import "testing"

func TestInterestingArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{[]int{9, 17}}, "Yes"},
		{"test2", args{[]int{1}}, "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InterestingArray(tt.args.A); got != tt.want {
				t.Errorf("InterestingArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkInterestingArray-8   	651772704	         1.931 ns/op	       0 B/op	       0 allocs/op
func BenchmarkInterestingArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InterestingArray([]int{9, 17})
	}
}
