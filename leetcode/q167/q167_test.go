package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	var testcases = []struct {
		input1   []int
		input2   int
		expect []int
	}{
		{[]int {2,7,11,15}, 9, []int{1,2}},
		{[]int {2,3,4}, 6, []int{1,3}},
		{[]int {-1,0}, -1, []int{1,2}},
	}
	for _, v := range testcases {
		output := twoSum(v.input1, v.input2)
		assert.Equal(t, v.expect, output)
	}
}
