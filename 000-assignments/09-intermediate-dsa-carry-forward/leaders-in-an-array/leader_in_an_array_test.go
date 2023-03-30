package leadersinanarray

import (
	"reflect"
	"testing"
)

func TestLeadersInAnArray(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][]int
	}{
		{"test1", args{[]int{16, 17, 4, 3, 5, 2}}, [][]int{{17, 5, 2}, {2, 5, 17}, {17, 2, 5}}},
		{"test2", args{[]int{1, 2, 3, 4, 0}}, [][]int{{4}}},
		{"test3", args{[]int{7, 4, 5, 7, 3}}, [][]int{{7, 3}, {3, 7}}},
		{"test_large", args{[]int{98, 23, 54, 12, 20, 7, 27}}, [][]int{{27, 54, 98}, {98, 54, 27}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got matches any of want result
			got := LeadersInAnArray(tt.args.A)
			var found bool
			for _, want := range tt.wantResult {
				if reflect.DeepEqual(got, want) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("LeadersInAnArray() = %v, want %v", got, tt.wantResult)
			}
		})
	}
}

// BenchmarkLeadersInAnArray-8   	12642969	        92.25 ns/op	      56 B/op	       3 allocs/op
func BenchmarkLeadersInAnArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LeadersInAnArray([]int{98, 23, 54, 12, 20, 7, 27})
	}
}
