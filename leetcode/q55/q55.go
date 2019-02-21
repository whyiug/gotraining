package main

/**
55. Jump Game
Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

Example 1:

Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum
             jump length is 0, which makes it impossible to reach the last index.
 */

import "fmt"

func main() {
	var nums55 = []int{3,0,8,2,0,0,1}
	fmt.Println(canJump(nums55))
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	var f = make([]int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 {
			if f[i-1] < i {
				break
			}
			if f[i-1] < i+nums[i] {
				f[i] = i + nums[i]
			} else {
				f[i] = f[i-1]
			}
		} else {
			f[i] = nums[i]
		}
	}
	return f[len(nums)-2] >= len(nums)-1
}

