package p01079

// Time Complexity: O(n^(n-1))
// Space Complexity: O(n)
func numTilePossibilities_backtracking(tiles string) int {
	memo := map[string]int{}
	return helper(memo, tiles, "") - 1
}

func helper(memo map[string]int, tiles, current string) int {
	cnt := 1
	for idx, ch := range tiles {
		next := current + string(ch)
		if memo[next] == 0 {
			if idx == 0 {
				cnt += helper(memo, tiles[1:], next)
			} else if idx == len(tiles)-1 {
				cnt += helper(memo, tiles[:idx], next)
			} else {
				cnt += helper(memo, tiles[:idx]+tiles[idx+1:], next)
			}
		}
	}

	memo[current] = cnt
	return cnt
}
