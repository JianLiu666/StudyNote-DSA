package p00823

import "sort"

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)

	memo := map[int]int{}
	for i, n := range arr {
		memo[n] = i
	}

	counter := make([]int, len(arr))
	for current, val := range arr {
		counter[current] = 1

		for left := 0; left < current; left++ {
			if val%arr[left] != 0 {
				continue
			}
			// 找到左邊節點，再來確定右邊節點的數字存不存在
			if right, exists := memo[val/arr[left]]; exists {
				counter[current] += mod(counter[left] * counter[right])
			}
		}
	}

	return mod(sum(counter))
}

func sum(arr []int) int {
	acc := 0
	for _, n := range arr {
		acc += n
	}

	return acc
}

func mod(n int) int {
	return n % 1000000007
}
