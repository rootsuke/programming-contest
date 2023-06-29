package main

func mergeAlternately(word1 string, word2 string) string {
	ans := ""

	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			ans += string(word1[i])
		}

		if i < len(word2) {
			ans += string(word2[i])
		}
	}

	return ans
}
