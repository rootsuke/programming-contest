package main

func solve(n int, dp []int) int {
	if n <= 2 {
		return n
	}

	if dp[n] != 0 {
		return dp[n]
	}

	dp[n] = solve(n-1, dp) + solve(n-2, dp)

	return dp[n]
}

func climbStairs(n int) int {
	dp := make([]int, n+1)

	return solve(n, dp)
}
