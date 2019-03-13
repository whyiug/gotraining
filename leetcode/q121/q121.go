package main

import "fmt"

func main() {
	// 121. Best Time to Buy and Sell Stock
	var ns = []int{7, 1, 5, 3, 6, 4}
	// var nums121 = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(ns))
	// 121. Best Time to Buy and Sell Stock
}

func maxProfit(prices []int) int {
	var minB, maxP int
	for i := 0; i < len(prices); i++ {
		if i > 0 {
			if prices[i]-minB > maxP {
				maxP = prices[i] - minB
			}
			if prices[i] < minB {
				minB = prices[i]
			}
		} else {
			minB = prices[i]
			maxP = 0
		}
	}
	return maxP
}
