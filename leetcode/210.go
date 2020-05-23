package main

import (
	"container/list"
	"fmt"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	result := []int{}
	inDegree := make([]int, numCourses)
	Map := make(map[int][]int, numCourses)
	for _, item := range prerequisites {
		inDegree[item[0]]++
		Map[item[1]] = append(Map[item[1]], item[0])
	}
	//for i := 0; i < numCourses; i++ {
	//	for j := 0; j < numCourses; j++ {
	//		if inDegree[j] == 0 {
	//			result = append(result, j)
	//			inDegree[j]--
	//			for _,k := range Map[j] {
	//				inDegree[k]--
	//			}
	//		}
	//	}
	//}
	queue := list.New()

	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue.PushBack(i)
		}
	}
	for queue.Len() > 0 {
		num := queue.Front()
		queue.Remove(num)
		result = append(result, num.Value.(int))
		for _, item := range Map[num.Value.(int)] {
			inDegree[item]--
			if inDegree[item] == 0 {
				queue.PushBack(item)
			}
		}
	}

	if len(result) == numCourses {
		return result
	}
	return []int{}
}

func main() {
	data := [][]int {{1,0}}
	num := 2
	fmt.Println(findOrder(num, data))
}
