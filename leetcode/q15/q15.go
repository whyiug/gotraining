package main

/**
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
Note:
The solution set must not contain duplicate triplets.
Example:
Given array nums = [-1, 0, 1, 2, -1, -4],
A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]

[0,0,0,0]
]
 */

import (
	"fmt"
	"reflect"
	"sort"
)

func main()  {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}
	nums := []int{0,0,0,0}
	fmt.Println(threeSum(nums))
}
func threeSum(nums []int) [][]int {
	result := [][]int{}
	numMap := map[int]int{}
	// 排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	// Hash
	for k, v := range nums {
		numMap[v] = k
	}
	// 两层遍历
	for k1 := 0; k1 < len(nums) - 1; k1 ++ {
		if k1 != 0 && nums[k1] == nums[k1-1] {
			continue
		}
		for k2 := k1+1; k2 < len(nums); k2 ++ {
			if k3, ok := numMap[-(nums[k1] + nums[k2])];ok && k3 > k2 && k3 > k1{
				if len(result) > 0 && sliceReflectEqual([]int{nums[k1], nums[k2], nums[k3]}, result[len(result)-1]) {
					continue
				} else {
					result = append(result, []int{nums[k1], nums[k2], -(nums[k1] + nums[k2])})
				}
			}
		}
	}
	return result
}

func sliceReflectEqual(a, b []int) bool {
	return reflect.DeepEqual(a, b)
}
