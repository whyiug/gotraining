package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	var testcases = []struct {
		input1   []int
		input2   int
		expect bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
	}
	for _, v := range testcases {
		output := canPlaceFlowers(v.input1, v.input2)
		assert.Equal(t, v.expect, output)
	}
}
