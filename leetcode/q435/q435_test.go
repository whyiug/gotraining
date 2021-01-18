package main

import (
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	var testcases = []struct {
		input   [][]int
		expect int
	}{
		{[][]int {{1,2},{2,3},{3,4},{1,3}}, 1},
	}
	for _, v := range testcases {
		output := eraseOverlapIntervals(v.input)
		if output != v.expect {
			t.Error()
		}
	}
}
