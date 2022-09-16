package p00017

// Time Complexity: O(m^n), where m are possible letters of each number, n is the length of digits
// Space Complexity: O(n)
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	mt := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := []string{}
	dfs(&res, mt, "", digits)
	return res
}

func dfs(res *[]string, mt map[byte]string, path, digits string) {
	if len(digits) == 1 {
		for _, ch := range mt[digits[0]] {
			*res = append(*res, path+string(ch))
		}
		return
	}

	for _, ch := range mt[digits[0]] {
		dfs(res, mt, path+string(ch), digits[1:])
	}
}
