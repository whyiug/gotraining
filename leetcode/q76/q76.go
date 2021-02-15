package main

import "fmt"

func main() {
	//fmt.Println(containString("abc", "acd"))
	//fmt.Println(containString("aa", "a")) // true
	//fmt.Println(containString("a", "aa")) // false
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "aa"))
}

//Input: s = "ADOBECODEBANC", t = "ABC"
//Output: "BANC"
func minWindow(s string, t string) string {
	var length, l, r int
	length = len(s)
	l = 0
	r = length - 1
	if containString(s[l:r+1], t) {
		for l < r {
			if containString(s[l+1:r+1], t) {
				l++
				continue
			} else if containString(s[l:r], t) {
				r--
				continue
			} else {
				break
			}
		}
		return s[l:r+1]
	} else {
		return ""
	}
}

// s是否包含t中所有字符
// aa 包含 a
// a 不包含 aa
func containString(s string, t string) bool {
	sMap := make(map[int32]int)
	for _, v := range s {
		sMap[v] ++
	}
	for _, v := range t {
		if _, exist := sMap[v]; exist && sMap[v] > 0 {
			sMap[v] --
		} else {
			return false
		}
	}
	return true
}
