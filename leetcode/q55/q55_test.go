package main

import (
	"testing"
)

func Test_canJump(t *testing.T) {
	var testcases = []struct {
		input  []int
		expect bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0, 2, 3}, false},
		{[]int{3, 0, 8, 2, 0, 0, 1}, true},
	}
	for _, v := range testcases {
		output := canJump(v.input)
		if output != v.expect {
			t.Error()
		}
	}
}
