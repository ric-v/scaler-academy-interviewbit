package tolower

import (
	"reflect"
	"testing"
)

func TestToLower(t *testing.T) {
	type args struct {
		A []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"test1", args{[]byte("Hello")}, []byte("hello")},
		{"test2", args{[]byte("HERE")}, []byte("here")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToLower(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkToLower-8   	267251418	         4.133 ns/op	       0 B/op	       0 allocs/op
func BenchmarkToLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToLower([]byte("Hello"))
	}
}
