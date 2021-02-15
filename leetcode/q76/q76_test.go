package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	var testcases = []struct {
		input1   string
		input2   string
		expect string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"cabwefgew cwaefgc f", "cae", "cwae"},
		//"aefgc"
	}
	for _, v := range testcases {
		assert.Equal(t, v.expect, minWindow(v.input1, v.input2))
	}
}
