package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	pos := make([]int, 60)
	target := make([]int, 60)
	minLen, minPos,maxPos := 0xffffffff,-1,-1
	for m := 0; m < len(t); m++ {
		target[int(t[m]) - 'A'] ++
	}
	judge := func() bool {
		for m := 0; m < len(t); m++ {
			if pos[int(t[m]) - 'A'] < target[int(t[m]) - 'A'] {
				return false
			}
		}
		return true
	}
	pos[int(s[0]) - 'A']++
	for i,j := 0,0; ; {
		if judge() == true {
			if minLen > j - i + 1 {
				minLen = j - i + 1
				minPos = i
				maxPos = j
			}
			tt := int(s[i]) - int('A')
			pos[tt]--
			i++
			if i == len(s) || j == len(s) {
				break
			}
		} else {
			j++
			if i == len(s) || j == len(s) {
				break
			}
			tt := int(s[j]) - int('A')
			pos[tt]++
		}
	}
	if minLen != 0xffffffff {
		result := ""
		for i := minPos; i <= maxPos; i++ {
			result = result + string(s[i])
		}
		return result
	}
	return ""
}

func main() {
	S := "a"
	T := "b"
	//for i := 0; i < len(T); i++ {
	//	fmt.Println(int(T[i]) - 'A')
	//}
	//fmt.Println(int('A') - 'A')
	fmt.Println(minWindow(S,T))
}