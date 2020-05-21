package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := []int{}
	Map := map[int]int{}
	for i, item := range nums {
		Map[item] = i
	}
	for key, value := range nums {
		newNum := target - value
		if mapVal,ok := Map[newNum]; ok == false || mapVal == key {
			continue
		}
		result = append(result, key, Map[newNum])
		return result
	}
	return result
}

func main() {
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
