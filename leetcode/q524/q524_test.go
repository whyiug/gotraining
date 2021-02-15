package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	var testcases = []struct {
		input struct {
			name string
			s    string
			d    []string
		}
		expect string
	}{
		{struct {
			name string
			s    string
			d    []string
		}{name: "case 1", s: "abpcplea", d: []string{
			"ale", "apple", "monkey", "plea",
		}}, "apple"},
		{struct {
			name string
			s    string
			d    []string
		}{name: "case 2", s: "abpcplea", d: []string{
			"a", "b", "c",
		}}, "a"},
	}
	for _, v := range testcases {
		assert.Equal(t, v.expect, findLongestWord(v.input.s, v.input.d))
	}
}
