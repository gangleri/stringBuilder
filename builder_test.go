package stringBuilder

import "testing"

func BenchmarkBuildString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buildString()
	}
}

func BenchmarkConcatString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatString()
	}
}

func BenchmarkAppendString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		appendString()
	}
}
