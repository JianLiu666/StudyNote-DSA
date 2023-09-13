package p00004

// Time Complexity: O(log(m+n)), where m is the number of nums1, n is the number of nums2
// Space Complexity: O(1)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)

	left := findKth(nums1, nums2, 0, n, 0, m, (n+m+1)/2)
	right := findKth(nums1, nums2, 0, n, 0, m, (n+m+2)/2)
	return float64(left+right) / 2
}

func findKth(nums1, nums2 []int, s1, e1, s2, e2, k int) int {
	l1, l2 := e1-s1, e2-s2
	if l1 > l2 {
		// 為了方便比較是否有其中一個 array 被排空了, 固定將第一個 array 設定成比較小的
		return findKth(nums2, nums1, s2, e2, s1, e1, k)
	}
	if l1 == 0 {
		return nums2[s2+k-1]
	}
	if k == 1 {
		return min(nums1[s1], nums2[s2])
	}

	// 確保不會溢位, 最多就取到 array 的最後一個元素
	i := s1 + min(l1, k/2) - 1
	j := s2 + min(l2, k/2) - 1

	if nums1[i] <= nums2[j] {
		return findKth(nums1, nums2, i+1, e1, s2, e2, k-(i-s1+1))
	} else {
		return findKth(nums1, nums2, s1, e1, j+1, e2, k-(j-s2+1))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
