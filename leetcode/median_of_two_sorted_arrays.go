package leetcode

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (m+n+1)/2 - partitionX

		maxLeftX := -1000001
		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}

		minRightX := 1000001
		if partitionX < m {
			minRightX = nums1[partitionX]
		}

		maxLeftY := -1000001
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}

		minRightY := 1000001
		if partitionY < n {
			minRightY = nums2[partitionY]
		}

		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (m+n)%2 != 0 {
				return float64(max(maxLeftX, maxLeftY))
			} else {
				return (float64(max(maxLeftX, maxLeftY)) + float64(min(minRightX, minRightY))) / 2.0
			}
		} else if maxLeftX > minRightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
