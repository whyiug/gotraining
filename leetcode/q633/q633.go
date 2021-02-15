package main

import (
	"math"
)

func main()  {
	judgeSquareSum(5)
}

//Input: c = 5
//Output: true
//Explanation: 1 * 1 + 2 * 2 = 5
func judgeSquareSum(c int) bool {
	if c == 0 {
		return true
	}
	var min, max int
	min = 0
	max = int(math.Floor(math.Sqrt(float64(c)))+1)
	for min <= max {
		sum := min*min + max*max
		if sum < c {
			min++
		} else if sum > c {
			max--
		} else {
			return true
		}
	}
	return false
}
