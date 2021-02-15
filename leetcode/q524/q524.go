package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	fmt.Println(int(math.Floor(math.Sqrt(4))))
	fmt.Println(int(math.Floor(math.Sqrt(8))))
	//fmt.Println(containThroughDeleting("ale", "abpcplea"))
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
}

//Input:
//s = "abpcplea", d = ["ale","apple","monkey","plea"]
//
//Output:
//"apple"
func findLongestWord(s string, d []string) string {
	var result []string
	for _, v := range d {
		if containThroughDeleting(v, s) {
			result = append(result, v)
		}
	}
	if len(result) == 0 {
		return ""
	} else {
		sort.SliceStable(result, func(i, j int) bool {
			if len(result[i]) > len(result[j]) {
				return true
			} else if len(result[i]) < len(result[j]) {
				return false
			} else {
				if strings.Compare(result[i], result[j]) < 0 {
					return true
				} else {
					return false
				}
			}
		})
		return result[0]
	}
}

// b 删除一些字母后是否生成a
func containThroughDeleting(a string, b string) bool {
	var aLength, bLength int
	aLength = len(a)
	bLength = len(b)
	var ai, bi int
	ai = 0
	bi = 0
	for ai < aLength && bi < bLength {
		if a[ai] == b[bi] {
			ai++
		}
		bi++
	}
	if ai == aLength {
		return true
	} else {
		return false
	}
	return false
}
