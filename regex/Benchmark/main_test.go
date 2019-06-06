package main

import (
	"strings"
	"testing"
)

func BenchmarkReplace2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strings.Replace("naufal naufal", "ziyad", "luthfiansyah", -1)
	}
}
func BenchmarkReplace5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strings.Replace("naufal naufal naufal naufal naufal",
			"ziyad", "lutfhfiansyah", -1)
	}
}
