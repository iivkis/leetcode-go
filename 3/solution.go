package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var up, down int
	var m = make(map[byte]int)
	var maxDistance int

	for down < len(s) {
		m[s[down]] += 1

		for m[s[down]] > 1 {
			m[s[up]] -= 1
			up += 1
		}

		down += 1

		distance := down - up
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
