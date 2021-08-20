package main

import "fmt"

func reverseStr(s string, k int) string {
	n := len(s)

	b := []byte(s)

	for i := 0; i < n; i += 2 * k {
		l := k
		if i+k > n {
			l = n - i
		}
		for j := 0; j < l/2; j++ {
			b[i+j], b[i+l-1-j] = b[i+l-1-j], b[i+j]
		}
	}

	return string(b)
}

func main() {
	for _, test := range []struct {
		s   string
		k   int
		exp string
	}{
		{"abcd", 2, "bacd"},
		{"abcde", 2, "bacde"},
		{"ab", 3, "ba"},
		{"abc", 3, "cba"},
		{"abcefghijk", 3, "cbaefgjihk"},
	} {
		r := reverseStr(test.s, test.k)
		fmt.Println(r, r == test.exp)
	}
}
