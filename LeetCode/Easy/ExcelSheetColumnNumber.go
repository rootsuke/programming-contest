package main

func titleToNumber(columnTitle string) int {
	ans := 0

	for _, char := range columnTitle {
		digit := char - 'A' + 1
		ans = ans*26 + int(digit)
	}

	return ans
}
