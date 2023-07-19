package sumofallsubmatrices

import "testing"

func TestSumOfAllSubmatrices(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[][]int{{1, 1}, {1, 1}}}, 16},
		{"test2", args{[][]int{{1, 2}, {3, 4}}}, 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfAllSubmatrices(tt.args.A); got != tt.want {
				t.Errorf("SumOfAllSubmatrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkSumOfAllSubmatrices-8   	148837458	         7.415 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSumOfAllSubmatrices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfAllSubmatrices([][]int{{1, 1}, {1, 1}})
	}
}
