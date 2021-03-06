package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	var testcases = []struct {
		input   [][]int
		expect int
	}{
		{[][]int {{1,2},{2,3},{3,4},{1,3}}, 1},
		{[][]int {{1,2},{1,2},{1,2}}, 2},
		{[][]int {{1,2},{2,3}}, 0},
	}
	for _, v := range testcases {
		output := eraseOverlapIntervals(v.input)
		assert.Equal(t, output, v.expect)
	}
}
