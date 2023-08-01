package main

func moveZeroes(nums []int) {
	slow := 0

	for fast := 0; fast < len(nums); fast++ {
		if slow == 0 && fast != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}

		if nums[slow] != 0 {
			slow++
		}
	}
}
