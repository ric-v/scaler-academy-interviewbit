package toupper

import (
	"reflect"
	"testing"
)

func TestToUpper(t *testing.T) {
	type args struct {
		A []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"test1", args{[]byte("abc def")}, []byte("ABC DEF")},
		{"test2", args{[]byte("abc def ghi")}, []byte("ABC DEF GHI")},
		{"test3", args{[]byte("abc def ghi jkl")}, []byte("ABC DEF GHI JKL")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUpper(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkToUpper-8   	100000000	        10.57 ns/op	       0 B/op	       0 allocs/op
func BenchmarkToUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToUpper([]byte("abc def ghi jkl"))
	}
}
