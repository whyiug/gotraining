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
		{[][]int {{10,16},{2,8},{1,6},{7,12}}, 2},
		{[][]int {{1,2},{3,4},{5,6},{7,8}}, 4},
		{[][]int {{1,2},{2,3},{3,4},{4,5}}, 2},
		{[][]int {{1,2}}, 1},
		{[][]int {{2,3},{2,3}}, 1},
	}
	for _, v := range testcases {
		output := findMinArrowShots(v.input)
		assert.Equal(t, v.expect, output, v.input)
	}
}
