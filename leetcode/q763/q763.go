package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("a"))
	fmt.Println(partitionLabels("eaaaabaaec"))
	fmt.Println(partitionLabels("eaaaabaaeccc"))
}
func partitionLabels(S string) []int {
	alphaMaxMap := make(map[string]int)
	sarr := strings.Split(S, "")
	for k, s := range sarr {
		if k >= alphaMaxMap[s] {
			alphaMaxMap[s] = k
		}
	}
	var result []int
	min, max := 0, 0
	len := len(S)
	for k, s := range sarr {
		if k < max {
			if alphaMaxMap[s] > max {
				max = alphaMaxMap[s]
			}
		} else if k == max {
			max = alphaMaxMap[s]
		} else {
			result = append(result, max-min+1)
			max = alphaMaxMap[s]
			min = k
		}
		if max == len -1 {
			result = append(result, max-min+1)
			break
		}
	}
	return result
}