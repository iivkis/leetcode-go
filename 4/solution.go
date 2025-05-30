package main

import "fmt"

// сложность n^2
func LongestPalindromeV2(s string) string {
	if len(s) < 2 {
		return s
	}

	search := func(up, down int) (int, int) {
		for ; up >= 0 && down < len(s); up, down = up-1, down+1 {
			if s[up] != s[down] {
				break
			}
		}
		return up + 1, down - 1
	}

	var maxUp, maxDown int
	var up, down int
	for i := range len(s) {
		up, down = search(i, i)
		if down-up > maxDown-maxUp {
			maxDown, maxUp = down, up
		}
		up, down = search(i, i+1)
		if down-up > maxDown-maxUp {
			maxDown, maxUp = down, up
		}
	}

	return s[maxUp : maxDown+1]
}

// сложность n^3
func LongestPalindromeV1(s string) string {
	var up, down int
	var maxUp, maxDown int

	isPalindrome := func(start int, end int) bool {
		for start < end {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
		return true
	}

	for ; up < len(s); up++ {
		for down = up; down < len(s); down++ {
			if isPalindrome(up, down) {
				if down-up > maxDown-maxUp {
					maxUp, maxDown = up, down
				}
			}
		}
	}

	return s[maxUp : maxDown+1]
}

func main() {
	result := LongestPalindromeV2("babad")
	fmt.Println(result)
}
