package main

func majorityElement(nums []int) int {
	dict := map[int]int{}
	max_key := 0

	for _, num := range nums {
		dict[num]++

		if dict[num] > len(nums)/2 {
			max_key = num
			break
		}
	}

	return max_key
}
