package main

import (
	"sort"
)

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) > len(dictionary[j]) {
			return true
		} else if len(dictionary[i]) == len(dictionary[j]) {
			return dictionary[i] < dictionary[j]
		}
		return false
	})

	for _, key := range dictionary {
		if match(s, key) {
			return key
		}
	}

	return ""
}

func match(s string, key string) bool {
	if len(s) < len(key) {
		return false
	}

	i := 0
	for j := 0; i < len(key) && j < len(s); j++ {
		if s[j] == key[i] {
			i++
		}
	}

	return i == len(key)
}
