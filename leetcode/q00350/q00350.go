package q00350

import "sort"

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func intersect_sort(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	ptr1 := 0
	ptr2 := 0

	result := []int{}
	for ptr1 < len(nums1) && ptr2 < len(nums2) {
		if nums1[ptr1] == nums2[ptr2] {
			result = append(result, nums1[ptr1])
			ptr1++
			ptr2++
		} else if nums1[ptr1] < nums2[ptr2] {
			ptr1++
		} else {
			ptr2++
		}
	}

	return result
}

// Time Complexity: O(n)
// Space Complexity: O(n)
func intersect_dict(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int, len(nums1))
	map2 := make(map[int]int, len(nums2))

	for _, num := range nums1 {
		map1[num]++
	}
	for _, num := range nums2 {
		map2[num]++
	}

	result := []int{}
	for m1_num, m1_count := range map1 {
		if m2_count, exists := map2[m1_num]; exists {
			count := m1_count
			if count > m2_count {
				count = m2_count
			}

			for i := 0; i < count; i++ {
				result = append(result, m1_num)
			}
		}
	}

	return result
}
