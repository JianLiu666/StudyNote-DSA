package p00804

import "strings"

// Time Complexity: O(n)
// Space Complexity: O(n)
func uniqueMorseRepresentations(words []string) int {
	memo := map[string]bool{}
	morse := map[rune]string{
		'a': ".-",
		'b': "-...",
		'c': "-.-.",
		'd': "-..",
		'e': ".",
		'f': "..-.",
		'g': "--.",
		'h': "....",
		'i': "..",
		'j': ".---",
		'k': "-.-",
		'l': ".-..",
		'm': "--",
		'n': "-.",
		'o': "---",
		'p': ".--.",
		'q': "--.-",
		'r': ".-.",
		's': "...",
		't': "-",
		'u': "..-",
		'v': "...-",
		'w': ".--",
		'x': "-..-",
		'y': "-.--",
		'z': "--..",
	}

	for _, word := range words {
		var str strings.Builder
		for _, ch := range word {
			str.WriteString(morse[ch])
		}
		memo[str.String()] = true
	}

	return len(memo)
}
