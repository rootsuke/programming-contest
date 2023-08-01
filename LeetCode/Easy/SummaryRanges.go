package main

import "fmt"

func summaryRanges(nums []int) []string {
	ans := []string{}

	if len(nums) == 0 {
		return ans
	}

	start := nums[0]

	if len(ans) == 1 {
		fmt.Println("run")
		ans = append(ans, fmt.Sprint(start))
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 != nums[i] {
			ans = append(ans, createRange(start, nums[i-1]))
			start = nums[i]
		}

		if i == len(nums)-1 {
			ans = append(ans, createRange(start, nums[i]))
		}
	}

	return ans
}

func createRange(start int, end int) string {
	if start == end {
		return fmt.Sprint(start)
	} else {
		return fmt.Sprintf("%v->%v", start, end)
	}
}
