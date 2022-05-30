package stringtoint

import "testing"

func BenchmarkStringToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100000; j++ {
			convert(j)
		}
	}
}
