// You are given an integer n representing the number of steps to reach the top of a staircase. You can climb with either 1 or 2 steps at a time.

// Return the number of distinct ways to climb to the top of the staircase.

package main

var cache = map[int]int{}

func main() {
	// Example usage
	n := 5 // Number of steps
	result := climbStairs(n)
	println(result) // Output: 8
}

func climbStairs(n int) (result int) {
	if cached, found := cache[n]; found {
		return cached
	}
	if n <= 2 {
		return n
	}
	cache[n] = climbStairs(n-1) + climbStairs(n-2)
	return cache[n]
}