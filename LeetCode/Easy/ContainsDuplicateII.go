package main

func containsNearbyDuplicate(nums []int, k int) bool {
	dict := make(map[int]int, 0)

	for i, v := range nums {
		if lastIdx, hasKey := dict[v]; hasKey && i-lastIdx <= k {
			return true
		}

		dict[v] = i
	}

	return false
}
