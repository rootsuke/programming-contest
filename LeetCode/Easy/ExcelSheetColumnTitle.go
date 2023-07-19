package main

func convertToTitle(n int) string {
	chars := []byte{}

	for n > 0 {
		chars = append(chars, 'A'+byte((n-1)%26))
		n = (n - 1) / 26
	}

	return string(reverse(chars))
}

func reverse(chars []byte) []byte {
	n := len(chars)
	for i := 0; i < n/2; i++ {
		chars[i], chars[n-i-1] = chars[n-i-1], chars[i]
	}

	return chars
}
