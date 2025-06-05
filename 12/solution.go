package main

import "fmt"

var L1 = map[int]string{
	1: "I",
	2: "II",
	3: "III",
	4: "IV",
	5: "V",
	6: "VI",
	7: "VII",
	8: "VIII",
	9: "IX",
}

var L2 = map[int]string{
	1: "X",
	2: "XX",
	3: "XXX",
	4: "XL",
	5: "L",
	6: "LX",
	7: "LXX",
	8: "LXXX",
	9: "XC",
}

var L3 = map[int]string{
	1: "C",
	2: "CC",
	3: "CCC",
	4: "CD",
	5: "D",
	6: "DC",
	7: "DCC",
	8: "DCCC",
	9: "CM",
}

var L4 = map[int]string{
	1: "M",
	2: "MM",
	3: "MMM",
}

var L = []map[int]string{L1, L2, L3, L4}

func intToRoman(num int) string {
	var ret string

	for i := 0; num > 0; i++ {
		x := num % 10
		num = num / 10
		ret = L[i][x] + ret
	}

	return ret
}

func main() {
	fmt.Println(intToRoman(3749))
}
