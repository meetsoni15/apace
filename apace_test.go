package apace

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	t.Parallel()
	type args struct {
		l       int
		charSet []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"TestRandomString 1", args{l: 13, charSet: nil}, 13, false},
		{"TestRandomString 2", args{l: 531, charSet: []string{"MMM"}}, 531, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandomString(tt.args.l, tt.args.charSet...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("Length of RandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRandomString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		RandomString(25)
	}
}

func TestRandomInt(t *testing.T) {
	tests := []struct {
		name string
		min  int
		max  int
	}{
		{
			name: "TestRandomInt",
			min:  12,
			max:  34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := RandomInt(tt.min, tt.max)
			if !(tt.min <= output && output <= tt.max) {
				t.Errorf("RandomInt() output is %v, want between %v and %v", output, tt.min, tt.max)

			}
		})
	}
}

func BenchmarkRandomInt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		RandomInt(0, 99999)
	}
}

func TestRandomArray(t *testing.T) {
	tests := []struct {
		name   string
		length int
		digits any
		want   int
	}{
		{"RandomArrayInt", 31, 99999, 31},
		{"RandomArrayString", 5500, "ABCD", 5500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomArray(tt.length, tt.digits); len(got) != tt.want {
				t.Errorf("RandomArray() got total elements are %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRandomArray(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		RandomArray(100, 99999)
	}
}
