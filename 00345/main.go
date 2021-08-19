package main

func reverseVowels(s string) string {
	b := []byte(s)
	for left, right := 0, len(b); left < right; {
		for left < right && !isVowels(b[left]) {
			left++
		}
		for left < right && !isVowels(b[right-1]) {
			right--
		}
		if left < right {
			b[left], b[right-1] = b[right-1], b[left]
			left++
			right--
		}
	}
	return string(b)
}

func isVowels(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
