package main

import "fmt"

func containsDuplicate(nums []int) bool {
	dict := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		if dict[nums[i]] > 0 {
			return true
		}

		fmt.Println(len(dict))

		dict[nums[i]]++
	}

	return false
}
