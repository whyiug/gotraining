package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	var testcases = []struct {
		input   string
		expect []int
	}{
		{"ababcbacadefegdehijhklij", []int{9,7,8}},
		{"a", []int{1}},
		{"eaaaabaaec", []int{9, 1}},
		{"eaaaabaaeccc", []int{9, 3}},
	}
	for _, v := range testcases {
		output := partitionLabels(v.input)
		assert.Equal(t, v.expect, output)
	}
}
