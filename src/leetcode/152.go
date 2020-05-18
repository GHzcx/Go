package main

import "fmt"

func maxProduct(nums []int) int {
	minMax := make([][]int, len(nums)) // 0 最小值 1最大值
	min := func(x, y int) int {if x < y {return x} else {return y}}
	max := func(x, y int) int {if x > y {return x} else {return y}}
	var result int
	for i, item := range nums {
		if i == 0 {
			minMax[i] = append(minMax[i], item)
			minMax[i] = append(minMax[i], item)
			result = minMax[i][1]
		} else {
			minMax[i] = append(minMax[i], min(item, min(item * minMax[i-1][0], item * minMax[i-1][1])))
			minMax[i] = append(minMax[i], max(item, max(item * minMax[i-1][0], item * minMax[i-1][1])))
			result = max(minMax[i][1], result)
		}
	}
	return result
}

func main() {
	nums := []int{2,3,-2,4}
	fmt.Println(maxProduct(nums))
}