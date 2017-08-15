package gramework

import "testing"

func TestBytesToString(t *testing.T) {
	tests := []struct {
		slice []byte
		want  string
	}{
		{
			slice: []byte("Hello world"),
			want:  "Hello world",
		},
		{
			slice: []byte("Hello world2304823049823049823409238402394823094823049428304"),
			want:  "Hello world2304823049823049823409238402394823094823049428304",
		},
	}
	for _, tt := range tests {
		if got := BytesToString(tt.slice); got != tt.want {
			t.Errorf("BytesToString() = %v, want %v", got, tt.want)
		}
	}
}

func BenchmarkBytesToString(b *testing.B) {
	s := []byte("Hello world2304823049823049823409238402394823094823049428304")
	b.ResetTimer()
	var a = ""
	for i := 0; i < b.N; i++ {
		a = BytesToString(s)
	}
	_ = a
}

func BenchmarkBytesToStringBuiltin(b *testing.B) {
	s := []byte("Hello world2304823049823049823409238402394823094823049428304")
	b.ResetTimer()
	var a = ""
	for i := 0; i < b.N; i++ {
		a = string(s)
	}
	_ = a
}
