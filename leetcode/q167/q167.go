package main

func main()  {
}

func twoSum(numbers []int, target int) []int {
	var length int
	length = len(numbers)
	var l, r int
	l = 0
	r = length - 1 - l
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum > target {
			r --
		} else if sum < target {
			l ++
		} else {
			return []int{l+1, r+1}
		}
	}
	return nil
}

