package main

import "fmt"

func validPalindrome(s string) bool {
	flag := 0
	len := len(s)
	ss := []rune(s)
	handle := func(mark, pos int) (int, int){
		if mark < 0 {
			return pos, len - 1 - pos + mark
		}
		return pos + mark, len - 1 - pos
	}
	judge := func(x,y int) bool{
		if x >= y {return true}
		if ss[x] == ss[y] {return true}
		return false
	}
	for i := 0 ; ; i++ {
		pos_x, pos_y := handle(flag, i)
		if pos_x >= pos_y {
			break
		}
		if ss[pos_x] == ss[pos_y] {
			continue
		} else if flag != 0 {
			return false
		} else if judge(pos_x+1, pos_y) &&  judge(pos_x + 1 + 1, pos_y - 1){
			flag = 1
			continue
		} else if judge(pos_x, pos_y - 1) && judge(pos_x + 1, pos_y - 1 - 1){
			flag = -1
			continue
		} else {
			return false
		}
	}
	return true
}

func main () {
	s := "aba"
	fmt.Println(validPalindrome(s))
}