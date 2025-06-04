package main

import "fmt"

// https://leetcode.com/problems/container-with-most-water/
// https://www.youtube.com/watch?v=UuiTKBwPgAo
func maxArea(height []int) int {
	var up, down int
	var maxS int

	down = len(height) - 1

	for up < down {
		if S := min(height[up], height[down]) * (down - up); S > maxS {
			maxS = S
		}

		if height[up] > height[down] {
			down -= 1
		} else {
			up += 1
		}
	}

	return maxS
}

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(arr))
}
