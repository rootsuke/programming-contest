package main

func maxProfit(prices []int) int {
	left := 0
	right := 1
	max := 0

	for right < len(prices) {
		if prices[left] < prices[right] {
			profit := prices[right] - prices[left]
			if max < profit {
				max = profit
			}
		} else {
			left = right
		}

		right++
	}

	return max
}
