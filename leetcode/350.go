package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	sort.Ints(nums1)
	sort.Ints(nums2)
	j := 0
	for i := 0; i < len(nums1); {
		if j >= len(nums2) {break}
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return res
}
