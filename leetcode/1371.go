package main

import "fmt"

func findTheLongestSubstring(s string) int {
	pos := func(s int) []int {
		arr := []int{0}
		for i := 1; i < 32; i++ {
			arr = append(arr, -1)
		}
		return arr
	}(32)
	newStr := []rune(s)
	status := 0
	MaxLength := 0
	for i := 0; i < len(newStr); i++ {
		if newStr[i] == 'a' {
			status ^= 1 << 0
		} else if newStr[i] == 'e' {
			status ^= 1 << 1
		} else if newStr[i] == 'i' {
			status ^= 1 << 2
		} else if newStr[i] == 'o' {
			status ^= 1 << 3
		} else if newStr[i] == 'u' {
			status ^= 1 << 4
		}
		if pos[status] != -1 {
			if i + 1 - pos[status] > MaxLength { MaxLength = i + 1 - pos[status]}
		} else {
			pos[status] = i + 1
		}
	}
	return MaxLength
}

func main() {
	s := "leetcodeisgreat"
	fmt.Println(findTheLongestSubstring(s))
}
