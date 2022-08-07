package p01220

// Time Complexity: O(n)
// Space Complexity: O(1)
func countVowelPermutation_optimize(n int) int {
	a, e, i, o, u := 1, 1, 1, 1, 1
	mod := 1000000007
	for idx := 2; idx <= n; idx++ {
		new_a := (e + i + u) % mod
		new_e := (a + i) % mod
		new_i := (e + o) % mod
		new_o := i % mod
		new_u := (i + o) % mod

		a, e, i, o, u = new_a, new_e, new_i, new_o, new_u
	}
	return (a + e + i + o + u) % mod
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func countVowelPermutation_array(n int) int {
	vowels := [5]int{1, 1, 1, 1, 1}
	rules := map[int][]int{
		0: {1},
		1: {0, 2},
		2: {0, 1, 3, 4},
		3: {2, 4},
		4: {0},
	}
	mod := 1000000007

	for i := 2; i <= n; i++ {
		new_vowels := [5]int{}
		for key, arr := range rules {
			for _, v := range arr {
				new_vowels[v] += vowels[key]
			}
		}
		for j := 0; j < 5; j++ {
			vowels[j] = new_vowels[j] % mod
		}
	}

	count := 0
	for _, val := range vowels {
		count += val
	}

	return count % mod
}
