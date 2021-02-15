package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	var testcases = []struct {
		input1   []int
		input2   int
		input3   []int
		input4   int
		expect []int
	}{
		{[]int {1,2,3,0,0,0}, 3, []int{2,5,6}, 3, []int{1,2,2,3,5,6}},
		{[]int {1}, 1, []int{}, 0, []int{1}},
		{[]int {0}, 0, []int{1}, 1, []int{1}},
	}
	for _, v := range testcases {
		merge(v.input1, v.input2, v.input3, v.input4)
		assert.Equal(t, v.expect, v.input1, "nage", v.input1)
	}
}
