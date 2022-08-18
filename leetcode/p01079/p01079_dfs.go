package p01079

func numTilePossibilities_dfs(tiles string) int {
	letters := [26]int{}
	for _, ch := range tiles {
		letters[ch-'A']++
	}

	return dfs(letters)
}

func dfs(letters [26]int) int {
	sum := 0
	for i := 0; i < 26; i++ {
		if letters[i] > 0 {
			sum++
			letters[i]--
			sum += dfs(letters)
			letters[i]++
		}
	}
	return sum
}
