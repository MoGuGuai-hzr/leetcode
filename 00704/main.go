package main

func search(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		middle := (i + j) >> 1
		if nums[middle] < target {
			i = middle + 1
		} else {
			j = middle
		}
	}
	if i == len(nums) || nums[i] != target {
		return -1
	}
	return i
}
