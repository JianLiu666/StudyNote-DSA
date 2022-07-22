package p00599

import "math"

// Time Complexity: O(n)
// Space Complexity: O(n)
func findRestaurant(list1 []string, list2 []string) []string {
	memo := map[string]int{}

	for idx, name := range list1 {
		memo[name] = idx
	}

	min := math.MaxInt
	res := []string{}
	for idx2, name := range list2 {
		if idx1, exists := memo[name]; exists {
			l := idx1 + idx2
			if l == min {
				res = append(res, name)
			} else if l < min {
				min = l
				res = []string{name}
			}
		}
	}

	return res
}
