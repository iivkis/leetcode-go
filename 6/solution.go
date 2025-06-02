package main

import (
	"fmt"
	"strings"
)

// O(n)
func convertV1(s string, numRows int) string {
	arr := make([][]byte, numRows)
	step := 1 // 1 or -1

	for i, j := 0, 0; j < len(s); i, j = i+step, j+1 {
		if step == -1 && i == 0 {
			step = 1
		} else if step == 1 && i == numRows-1 {
			step = -1
		}
		arr[i] = append(arr[i], s[j])
	}

	strbld := strings.Builder{}
	for _, b := range arr {
		strbld.Write(b)
	}
	return strbld.String()
}

func main() {
	result := convertV1("PAYPALISHIRING", 3)
	fmt.Println(result)
}
