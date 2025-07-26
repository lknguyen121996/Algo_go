package main

import (
	"fmt"
	"sort"
)

// You are given an integer array nums and an integer k. You may partition nums into one or more subsequences such that each element in nums appears in exactly one of the subsequences.

// Return the minimum number of subsequences needed such that the difference between the maximum and minimum values in each subsequence is at most k.

// A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.
func main() {
	type test struct {
		nums   []int
		k      int
		result int
	}
	tests := []test{
		{
			nums:   []int{3, 6, 1, 2, 5},
			k:      2,
			result: 2,
		},
		{
			nums:   []int{1, 2, 3},
			k:      1,
			result: 2,
		},
		{
			nums:   []int{2, 2, 4, 5},
			k:      0,
			result: 3,
		},
		{
			nums:   []int{3, 1, 3, 4, 2},
			k:      0,
			result: 4,
		},
	}
	for _, test := range tests {
		ok := test.result == partitionArray(test.nums, test.k)
		fmt.Println(ok)
		if !ok {
			fmt.Printf("Expected: %d, Result:%d", test.result, partitionArray(test.nums, test.k))
		}
	}
}

func partitionArray(nums []int, k int) int {
	result := make([][]int, 0)
	sort.Ints(nums)
	max := nums[0]
	tmpSlice := make([]int, 0)
	for _, num := range nums {
		//  check the num and add to new slice
		if len(tmpSlice) == 0 || num <= max+k {
			tmpSlice = append(tmpSlice, num)
			continue
		}
		result = append(result, tmpSlice)
		tmpSlice = make([]int, 0)
		tmpSlice = append(tmpSlice, num)
		max = num
	}
	if len(tmpSlice) > 0 {
		result = append(result, tmpSlice)
	}
	return len(result)
}
