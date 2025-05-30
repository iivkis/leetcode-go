// go test -bench=. -benchmem
// goos: linux
// goarch: amd64
// pkg: leetcode-go/4
// cpu: AMD Ryzen 5 5500
// BenchmarkLongestPalindromeV1-12           117408             10190 ns/op               0 B/op          0 allocs/op
// BenchmarkLongestPalindromeV2-12          3142077               379.5 ns/op             0 B/op          0 allocs/op
// PASS
// ok      leetcode-go/4   2.393s

package main

import "testing"

const TEST_STRING = "nckahshkchalcnkzcuhaladczujygsdbabwertyujsxcvbzxcvbndfghretyuxcvbndfghjcbzhjdhihufbdhfvakudhuandhdbfgasjkafdibfalfvuagfhfladasdaDaf"

func BenchmarkLongestPalindromeV1(b *testing.B) {
	for b.Loop() {
		LongestPalindromeV1(TEST_STRING)
	}
}

func BenchmarkLongestPalindromeV2(b *testing.B) {
	for b.Loop() {
		LongestPalindromeV2(TEST_STRING)
	}
}
