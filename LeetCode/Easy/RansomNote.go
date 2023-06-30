package main

func canConstruct(ransomNote string, magazine string) bool {
	charCount := make(map[byte]int, 0)
	for i := 0; i < len(magazine); i++ {
		char := magazine[i]
		charCount[char]++
	}

	for i := 0; i < len(ransomNote); i++ {
		if charCount[ransomNote[i]] == 0 {
			return false
		}

		charCount[ransomNote[i]]--
	}

	return true
}
