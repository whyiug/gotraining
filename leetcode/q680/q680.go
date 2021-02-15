package main

import "fmt"

func main() {
	fmt.Println(validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"))
}

func validPalindrome(s string) bool {
	var length int
	length = len(s)
	var l, r int
	r = 0
	r = length - 1
	for l <= r {
		if s[l] == s[r] {
			l++
			r--
			continue
		} else {
			return  isPalindrome(s[l+1:r+1]) || isPalindrome(s[l:r])
			//return (l+1 < length && isPalindrome(s[l+1:r+1])) || (r >= 0 && isPalindrome(s[l:r]))
		}
	}
	return true
}

// 判断字符串是否回文
func isPalindrome(s string) bool {
	var length int
	length = len(s)
	var l, r int
	r = 0
	r = length - 1
	for l <= r {
		if s[l] == s[r] {
			l++
			r--
			continue
		} else {
			return false
		}
	}
	return true
}
