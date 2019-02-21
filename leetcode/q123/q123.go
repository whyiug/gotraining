package main

import "fmt"

func main() {
	// var nums123 = []int{1,2,3,4,5}
	var nums123 = []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	// var nums123 = []int{7,6,4,3,1}
	// var nums123 = []int{6,1,3,2,4,7}
	// var nums123 = []int{3,3,5,0,0,3,1,4}
	fmt.Println(maxProfit(nums123))
}

// 123. Best Time to Buy and Sell Stock III
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	local := make([]int, 3)
	global := make([]int, 3)
	var diff int
	for i := 0; i < len(prices)-1; i++ {
		diff = prices[i+1] - prices[i]
		for j := 2; j >= 1; j-- {
			local[j] = GetMax(global[j-1]+GetMax(diff, 0), local[j]+diff)
			global[j] = GetMax(local[j], global[j])
		}
	}
	return global[2]
}

// 123. Best Time to Buy and Sell Stock III
// WA
func maxProfit123WA(prices []int) int {
	var begin, end int
	var r []int
	begin = 0
	end = 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			end = i
			if end == len(prices)-1 {
				r = append(r, prices[end]-prices[begin])
			}
		} else {
			if end-begin >= 1 {
				r = append(r, prices[end]-prices[begin])
			}
			begin = i
			end = i
		}
	}
	fmt.Println(r)
	max1 := 0
	max2 := 0
	for _, v := range r {
		if v > max1 {
			max2, max1 = max1, v
		} else if v > max2 {
			max2 = v
		}
	}
	return max1 + max2
}

func GetMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
