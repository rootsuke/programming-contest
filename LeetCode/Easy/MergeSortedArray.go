package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1

	for k := 1; k <= m+n; k++ {
		if i < 0 {
			nums1[m+n-k] = nums2[j]
			j--
			continue
		}

		if j < 0 {
			nums1[m+n-k] = nums1[i]
			i--
			continue
		}

		if nums1[i] >= nums2[j] {
			nums1[m+n-k] = nums1[i]
			i--
		} else {
			nums1[m+n-k] = nums2[j]
			j--
		}
	}
}
