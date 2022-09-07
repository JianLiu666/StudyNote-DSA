package p00997

// Time Complexity: O(m), where m is the number of trust
// Space Complexity: O(n), where n is the number of people
func findJudge_hash(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}

	// 紀錄 people[i] 是否相信過人
	memo := map[int]bool{}
	// 紀錄 people[i] 的得票數
	votes := map[int]int{}
	for _, pair := range trust {
		memo[pair[0]] = true
		votes[pair[1]]++
	}

	for person, votes := range votes {
		// 不曾相信過人且受到所有人的信任
		if _, exists := memo[person]; !exists && votes == n-1 {
			return person
		}
	}

	return -1
}

// Time Complexity: O(m), where m is the number of trust
// Space Complexity: O(n), where n is the number of people
func findJudge_graph(n int, trust [][]int) int {
	degree := make([]int, n+1)
	for _, pair := range trust {
		degree[pair[0]]--
		degree[pair[1]]++
	}

	for i := 1; i <= n; i++ {
		if degree[i] == n-1 {
			return i
		}
	}

	return -1
}
