package main

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		intersect(nums2, nums1)
	}

	hashTable := toHashTable(nums1)
	ans := []int{}

	for _, v := range nums2 {
		if hashTable[v] > 0 {
			ans = append(ans, v)
			hashTable[v]--
		}
	}

	return ans
}

func toHashTable(nums []int) map[int]int {
	hash := make(map[int]int, 0)
	for _, v := range nums {
		hash[v]++
	}

	return hash
}
