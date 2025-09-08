package main

import "fmt"

func main() {
	type Test struct {
		s      string
		result string
	}
	tests := []Test{
		{"abc", "cba"},
		{"b", "b"},
		{"axcd", "dcxa"},
	}
	for _, test := range tests {
		if result := reverseString([]byte(test.s)); result != test.result {
			fmt.Printf("Input: s: %v", test.s)
			fmt.Printf("Result: %v \n", result)
			fmt.Printf("Expected Result: %v \n", test.result)
		}
		if result := reverseString([]byte(test.s)); result != test.result {
			fmt.Printf("Input: s: %v", test.s)
			fmt.Printf("Result: %v \n", result)
			fmt.Printf("Expected Result: %v \n", test.result)
		}
	}
}

func reverseString(s []byte) string {
	i := 0
	j := len(s) - 1
	for i < j {
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
		i++
		j--
	}
	return string(s)
}
