package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_validPalindrome(t *testing.T) {
	var testcases = []struct {
		input   string
		expect bool
	}{
		{"aba", true},
		{"abca", true},
		{"abbca", true},
		{"acbba", true},
		{"a", true},
		{"ab", true},
		{"abc", false},
		{"aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga", true},
	}
	for _, v := range testcases {
		output := validPalindrome(v.input)
		assert.Equal(t, v.expect, output)
	}
}
