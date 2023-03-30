package sumofallsubarrays

import "testing"

func TestSumOfAllSubarrays(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test1", args{[]int{1, 2, 3}}, 20},
		{"test2", args{[]int{2, 1, 3}}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfAllSubarrays(tt.args.arr); got != tt.want {
				t.Errorf("SumOfAllSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSumOfAllSubarrays-8   	548010015	         2.082 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSumOfAllSubarrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfAllSubarrays([]int{1, 2, 3})
	}
}

// BenchmarkSumOfAllSubarrays_CarryForward-8   	143303241	         8.472 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSumOfAllSubarrays_CarryForward(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfAllSubarrays_CarryForward([]int{1, 2, 3})
	}
}

// BenchmarkSumOfAllSubarrays_PrefixSum-8   	13045264	        89.61 ns/op	      48 B/op	       2 allocs/op
func BenchmarkSumOfAllSubarrays_PrefixSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfAllSubarrays_PrefixSum([]int{1, 2, 3})
	}
}

// BenchmarkSumOfAllSubarrays_BruteForce-8   	71949778	        17.10 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSumOfAllSubarrays_BruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfAllSubarrays_BruteForce([]int{1, 2, 3})
	}
}
