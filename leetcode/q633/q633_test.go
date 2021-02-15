package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	var testcases = []struct {
		input   int
		expect bool
	}{
		{5, true},
		{3, false},
		{4, true},
		{2, true},
		{1, true},
		{0, true},
	}
	for _, v := range testcases {
		assert.Equal(t, v.expect, judgeSquareSum(v.input))
	}
}
