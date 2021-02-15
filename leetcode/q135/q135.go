package main

import "fmt"

func main() {
	fmt.Println(candy([]int{1, 3, 2, 2, 1}))
	fmt.Println(candy([]int{1, 0, 2}))
	fmt.Println(candy([]int{1, 2, 2}))
}

func candy(ratings []int) int {
	minValue := 1
	if len(ratings) == 1 {
		return minValue
	}
	len := len(ratings)
	results := make([]int, len)
	results[0] = minValue
	for k := 1; k < len; k++ {
		if ratings[k] > ratings[k-1] {
			results[k] = results[k-1] + 1
		} else {
			results[k] = minValue
		}
	}
	for k := len - 2; k >= 0; k-- {
		if ratings[k] > ratings[k+1] && results[k] <= results[k+1] {
			results[k] = results[k+1] + 1
		}
	}
	sum := 0
	for k := 0; k < len; k++ {
		sum += results[k]
	}
	return sum
}
