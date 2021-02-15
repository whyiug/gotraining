package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	var testcases = []struct {
		input   []int
		expect int
	}{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
		{[]int{1,3,2,2,1}, 7},
		{[]int{1,3,4,5,2}, 11},
	}
	for _, v := range testcases {
		output := candy(v.input)
		assert.Equal(t, v.expect, output, v.input)
	}
}
