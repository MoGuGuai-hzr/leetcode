package main

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	cnt := 0
	for left, right := 0, len(people); left < right; {
		if people[left]+people[right-1] <= limit {
			left++
			right--
		} else {
			right--
		}
		cnt++
	}

	return cnt
}
