package main

import "fmt"

/*
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none)
of the characters without disturbing the relative positions of the remaining characters.
(i.e., "ace" is a subsequence of "abcde" while "aec" is not).
*/
func IsSubsequence(s, t string) bool {
	result := true
	j := 0
	if len(s) == 0 {
		return result
	}
	for i := 0; i < len(t); i++ {
		if j > len(s)-1 {
			break
		}
		if s[j] == t[i] {
			j++
		}
	}
	if j < len(s) {
		result = false
	}
	return result
}

func IsSubsequenceEnforce(s, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

func main() {
	type Test struct {
		s      string
		t      string
		result bool
	}
	tests := []Test{
		{"abc", "ahbgdc", true},
		{"b", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}
	for _, test := range tests {
		if result := IsSubsequence(test.s, test.t); result != test.result {
			fmt.Printf("Input: s: %v, t: %v \n", test.s, test.t)
			fmt.Printf("Result: %v \n", result)
			fmt.Printf("Expected Result: %v \n", test.result)
		}
		if result := IsSubsequenceEnforce(test.s, test.t); result != test.result {
			fmt.Printf("Input: s: %v, t: %v \n", test.s, test.t)
			fmt.Printf("Result: %v \n", result)
			fmt.Printf("Expected Result: %v \n", test.result)
		}
	}
}
